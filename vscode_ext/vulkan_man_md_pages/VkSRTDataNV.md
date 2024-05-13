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

S=⎝⎜⎛​sx00​asy0​bcsz​pvxpvypvz​⎠⎟⎞​

The rotation quaternion is defined as:

  
`R` = \[ `qx`, `qy`, `qz`, `qw` \]

This is a rotation around a conceptual normalized axis \[ ax, ay, az \]
of amount `theta` such that:

  
\[ `qx`, `qy`, `qz` \] = sin(`theta`/2) × \[ `ax`, `ay`, `az` \]

and

  
`qw` = cos(`theta`/2)

Finally, the transform has a translation T constructed from the
parameters above by:

T=⎝⎜⎛​100​010​001​txtytz​⎠⎟⎞​

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

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
