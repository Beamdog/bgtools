bammer.exe is used to dissemble BAMs to PNGs and (re)assemble PNGs to BAMs, using the BAMD data to build the BAMs. Typical syntax is something like (from a command prompt, .bat file or WeiDU): 

To convert BAM to PNGs (with BAMD output):
bammer -input=mybams/mybam.bam -output=mybams/mybam > mybams/mybam.bamd

To convert PNGs to BAMs (using BAMD input):
bammer -input=mybams/mybam.bam [-mirror=true] [-offsetx=5] [-offsety=-5] > mybams/mybam.bamd

  Optional parameters (BAM INPUT):
  -mirror:  flips the animation on the vertical axis if true
  -offsetx: moves all BAM frames horizontally by the specified number of pixels (+/- the existing offset)
  -offsety: moves all BAM frames vertically by the specified number of pixels (+/- the existing offset)

  Optional parameters (BAM OUTPUT):
  -palette: set a replacement paletted png
