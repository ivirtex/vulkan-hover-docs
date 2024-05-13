# VkViewportSwizzleNV(3) Manual Page

## Name

VkViewportSwizzleNV - Structure specifying a viewport swizzle



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkViewportSwizzleNV` structure is defined as:

``` c
// Provided by VK_NV_viewport_swizzle
typedef struct VkViewportSwizzleNV {
    VkViewportCoordinateSwizzleNV    x;
    VkViewportCoordinateSwizzleNV    y;
    VkViewportCoordinateSwizzleNV    z;
    VkViewportCoordinateSwizzleNV    w;
} VkViewportSwizzleNV;
```

## <a href="#_members" class="anchor"></a>Members

- `x` is a
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value specifying the swizzle operation to apply to the x component of
  the primitive

- `y` is a
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value specifying the swizzle operation to apply to the y component of
  the primitive

- `z` is a
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value specifying the swizzle operation to apply to the z component of
  the primitive

- `w` is a
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value specifying the swizzle operation to apply to the w component of
  the primitive

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkViewportSwizzleNV-x-parameter"
  id="VUID-VkViewportSwizzleNV-x-parameter"></a>
  VUID-VkViewportSwizzleNV-x-parameter  
  `x` **must** be a valid
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value

- <a href="#VUID-VkViewportSwizzleNV-y-parameter"
  id="VUID-VkViewportSwizzleNV-y-parameter"></a>
  VUID-VkViewportSwizzleNV-y-parameter  
  `y` **must** be a valid
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value

- <a href="#VUID-VkViewportSwizzleNV-z-parameter"
  id="VUID-VkViewportSwizzleNV-z-parameter"></a>
  VUID-VkViewportSwizzleNV-z-parameter  
  `z` **must** be a valid
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value

- <a href="#VUID-VkViewportSwizzleNV-w-parameter"
  id="VUID-VkViewportSwizzleNV-w-parameter"></a>
  VUID-VkViewportSwizzleNV-w-parameter  
  `w` **must** be a valid
  [VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html),
[VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html),
[VkViewportCoordinateSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportCoordinateSwizzleNV.html),
[vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkViewportSwizzleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
