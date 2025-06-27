# VkSRTDataNV(3) Manual Page

## Name

VkSRTDataNV - Structure specifying a transform in SRT decomposition



## <a href="#_c_specification" class="anchor"></a>C Specification

An acceleration structure SRT transform is defined by the structure:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkSRTDataNV {
    float    sx;
    float    a;
    float    b;
    float    pvx;
    float    sy;
    float    c;
    float    pvy;
    float    sz;
    float    pvz;
    float    qx;
    float    qy;
    float    qz;
    float    qw;
    float    tx;
    float    ty;
    float    tz;
} VkSRTDataNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sx` is the x component of the scale of the transform

- `a` is one component of the shear for the transform

- `b` is one component of the shear for the transform

- `pvx` is the x component of the pivot point of the transform

- `sy` is the y component of the scale of the transform

- `c` is one component of the shear for the transform

- `pvy` is the y component of the pivot point of the transform

- `sz` is the z component of the scale of the transform

- `pvz` is the z component of the pivot point of the transform

- `qx` is the x component of the rotation quaternion

- `qy` is the y component of the rotation quaternion

- `qz` is the z component of the rotation quaternion

- `qw` is the w component of the rotation quaternion

- `tx` is the x component of the post-rotation translation

- `ty` is the y component of the post-rotation translation

- `tz` is the z component of the post-rotation translation

## <a href="#_description" class="anchor"></a>Description

This transform decomposition consists of three elements. The first is a
matrix S, consisting of a scale, shear, and translation, usually used to
define the pivot point of the following rotation. This matrix is
constructed from the parameters above by:

S=![](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIwLjg3NWVtIiBoZWlnaHQ9IjMuNjAwZW0iIHZpZXdib3g9IjAgMCA4NzUgMzYwMCI+PHBhdGggZD0iTTg2Myw5YzAsLTIsLTIsLTUsLTYsLTljMCwwLC0xNywwLC0xNywwYy0xMi43LDAsLTE5LjMsMC4zLC0yMCwxCmMtNS4zLDUuMywtMTAuMywxMSwtMTUsMTdjLTI0Mi43LDI5NC43LC0zOTUuMyw2ODIsLTQ1OCwxMTYyYy0yMS4zLDE2My4zLC0zMy4zLDM0OSwKLTM2LDU1NyBsMCw4NGMwLjIsNiwwLDI2LDAsNjBjMiwxNTkuMywxMCwzMTAuNywyNCw0NTRjNTMuMyw1MjgsMjEwLAo5NDkuNyw0NzAsMTI2NWM0LjcsNiw5LjcsMTEuNywxNSwxN2MwLjcsMC43LDcsMSwxOSwxYzAsMCwxOCwwLDE4LDBjNCwtNCw2LC03LDYsLTkKYzAsLTIuNywtMy4zLC04LjcsLTEwLC0xOGMtMTM1LjMsLTE5Mi43LC0yMzUuNSwtNDE0LjMsLTMwMC41LC02NjVjLTY1LC0yNTAuNywtMTAyLjUsCi01NDQuNywtMTEyLjUsLTg4MmMtMiwtMTA0LC0zLC0xNjcsLTMsLTE4OQpsMCwtOTJjMCwtMTYyLjcsNS43LC0zMTQsMTcsLTQ1NGMyMC43LC0yNzIsNjMuNywtNTEzLDEyOSwtNzIzYzY1LjMsCi0yMTAsMTU1LjMsLTM5Ni4zLDI3MCwtNTU5YzYuNywtOS4zLDEwLC0xNS4zLDEwLC0xOHoiIC8+PC9zdmc+)​sx00​asy0​bcsz​pvxpvypvz​![](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIwLjg3NWVtIiBoZWlnaHQ9IjMuNjAwZW0iIHZpZXdib3g9IjAgMCA4NzUgMzYwMCI+PHBhdGggZD0iTTc2LDBjLTE2LjcsMCwtMjUsMywtMjUsOWMwLDIsMiw2LjMsNiwxM2MyMS4zLDI4LjcsNDIuMyw2MC4zLAo2Myw5NWM5Ni43LDE1Ni43LDE3Mi44LDMzMi41LDIyOC41LDUyNy41YzU1LjcsMTk1LDkyLjgsNDE2LjUsMTExLjUsNjY0LjUKYzExLjMsMTM5LjMsMTcsMjkwLjcsMTcsNDU0YzAsMjgsMS43LDQzLDMuMyw0NWwwLDkKYy0zLDQsLTMuMywxNi43LC0zLjMsMzhjMCwxNjIsLTUuNywzMTMuNywtMTcsNDU1Yy0xOC43LDI0OCwtNTUuOCw0NjkuMywtMTExLjUsNjY0CmMtNTUuNywxOTQuNywtMTMxLjgsMzcwLjMsLTIyOC41LDUyN2MtMjAuNywzNC43LC00MS43LDY2LjMsLTYzLDk1Yy0yLDMuMywtNCw3LC02LDExCmMwLDcuMyw1LjcsMTEsMTcsMTFjMCwwLDExLDAsMTEsMGM5LjMsMCwxNC4zLC0wLjMsMTUsLTFjNS4zLC01LjMsMTAuMywtMTEsMTUsLTE3CmMyNDIuNywtMjk0LjcsMzk1LjMsLTY4MS43LDQ1OCwtMTE2MWMyMS4zLC0xNjQuNywzMy4zLC0zNTAuNywzNiwtNTU4CmwwLC0xNDRjLTIsLTE1OS4zLC0xMCwtMzEwLjcsLTI0LC00NTRjLTUzLjMsLTUyOCwtMjEwLC05NDkuNywKLTQ3MCwtMTI2NWMtNC43LC02LC05LjcsLTExLjcsLTE1LC0xN2MtMC43LC0wLjcsLTYuNywtMSwtMTgsLTF6IiAvPjwvc3ZnPg==)​

The rotation quaternion is defined as:

  
`R` = \[ `qx`, `qy`, `qz`, `qw` \]

This is a rotation around a conceptual normalized axis \[ ax, ay, az \]
of amount `theta` such that:

  
\[ `qx`, `qy`, `qz` \] = sin(`theta`/2) × \[ `ax`, `ay`, `az` \]

and

  
`qw` = cos(`theta`/2)

Finally, the transform has a translation T constructed from the
parameters above by:

T=![](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIwLjg3NWVtIiBoZWlnaHQ9IjMuNjAwZW0iIHZpZXdib3g9IjAgMCA4NzUgMzYwMCI+PHBhdGggZD0iTTg2Myw5YzAsLTIsLTIsLTUsLTYsLTljMCwwLC0xNywwLC0xNywwYy0xMi43LDAsLTE5LjMsMC4zLC0yMCwxCmMtNS4zLDUuMywtMTAuMywxMSwtMTUsMTdjLTI0Mi43LDI5NC43LC0zOTUuMyw2ODIsLTQ1OCwxMTYyYy0yMS4zLDE2My4zLC0zMy4zLDM0OSwKLTM2LDU1NyBsMCw4NGMwLjIsNiwwLDI2LDAsNjBjMiwxNTkuMywxMCwzMTAuNywyNCw0NTRjNTMuMyw1MjgsMjEwLAo5NDkuNyw0NzAsMTI2NWM0LjcsNiw5LjcsMTEuNywxNSwxN2MwLjcsMC43LDcsMSwxOSwxYzAsMCwxOCwwLDE4LDBjNCwtNCw2LC03LDYsLTkKYzAsLTIuNywtMy4zLC04LjcsLTEwLC0xOGMtMTM1LjMsLTE5Mi43LC0yMzUuNSwtNDE0LjMsLTMwMC41LC02NjVjLTY1LC0yNTAuNywtMTAyLjUsCi01NDQuNywtMTEyLjUsLTg4MmMtMiwtMTA0LC0zLC0xNjcsLTMsLTE4OQpsMCwtOTJjMCwtMTYyLjcsNS43LC0zMTQsMTcsLTQ1NGMyMC43LC0yNzIsNjMuNywtNTEzLDEyOSwtNzIzYzY1LjMsCi0yMTAsMTU1LjMsLTM5Ni4zLDI3MCwtNTU5YzYuNywtOS4zLDEwLC0xNS4zLDEwLC0xOHoiIC8+PC9zdmc+)​100​010​001​txtytz​![](data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIwLjg3NWVtIiBoZWlnaHQ9IjMuNjAwZW0iIHZpZXdib3g9IjAgMCA4NzUgMzYwMCI+PHBhdGggZD0iTTc2LDBjLTE2LjcsMCwtMjUsMywtMjUsOWMwLDIsMiw2LjMsNiwxM2MyMS4zLDI4LjcsNDIuMyw2MC4zLAo2Myw5NWM5Ni43LDE1Ni43LDE3Mi44LDMzMi41LDIyOC41LDUyNy41YzU1LjcsMTk1LDkyLjgsNDE2LjUsMTExLjUsNjY0LjUKYzExLjMsMTM5LjMsMTcsMjkwLjcsMTcsNDU0YzAsMjgsMS43LDQzLDMuMyw0NWwwLDkKYy0zLDQsLTMuMywxNi43LC0zLjMsMzhjMCwxNjIsLTUuNywzMTMuNywtMTcsNDU1Yy0xOC43LDI0OCwtNTUuOCw0NjkuMywtMTExLjUsNjY0CmMtNTUuNywxOTQuNywtMTMxLjgsMzcwLjMsLTIyOC41LDUyN2MtMjAuNywzNC43LC00MS43LDY2LjMsLTYzLDk1Yy0yLDMuMywtNCw3LC02LDExCmMwLDcuMyw1LjcsMTEsMTcsMTFjMCwwLDExLDAsMTEsMGM5LjMsMCwxNC4zLC0wLjMsMTUsLTFjNS4zLC01LjMsMTAuMywtMTEsMTUsLTE3CmMyNDIuNywtMjk0LjcsMzk1LjMsLTY4MS43LDQ1OCwtMTE2MWMyMS4zLC0xNjQuNywzMy4zLC0zNTAuNywzNiwtNTU4CmwwLC0xNDRjLTIsLTE1OS4zLC0xMCwtMzEwLjcsLTI0LC00NTRjLTUzLjMsLTUyOCwtMjEwLC05NDkuNywKLTQ3MCwtMTI2NWMtNC43LC02LC05LjcsLTExLjcsLTE1LC0xN2MtMC43LC0wLjcsLTYuNywtMSwtMTgsLTF6IiAvPjwvc3ZnPg==)​

The effective derived transform is then given by

  
`T` × `R` × `S`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSRTDataNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
