# VkComponentSwizzle(3) Manual Page

## Name

VkComponentSwizzle - Specify how a component is swizzled



## [](#_c_specification)C Specification

Possible values of the members of [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), specifying the component values placed in each component of the output vector, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkComponentSwizzle {
    VK_COMPONENT_SWIZZLE_IDENTITY = 0,
    VK_COMPONENT_SWIZZLE_ZERO = 1,
    VK_COMPONENT_SWIZZLE_ONE = 2,
    VK_COMPONENT_SWIZZLE_R = 3,
    VK_COMPONENT_SWIZZLE_G = 4,
    VK_COMPONENT_SWIZZLE_B = 5,
    VK_COMPONENT_SWIZZLE_A = 6,
} VkComponentSwizzle;
```

## [](#_description)Description

- `VK_COMPONENT_SWIZZLE_IDENTITY` specifies that the component is set to the identity swizzle.
- `VK_COMPONENT_SWIZZLE_ZERO` specifies that the component is set to zero.
- `VK_COMPONENT_SWIZZLE_ONE` specifies that the component is set to either 1 or 1.0, depending on whether the type of the image view format is integer or floating-point respectively, as determined by the [Format Definition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-definition) section for each [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html).
- `VK_COMPONENT_SWIZZLE_R` specifies that the component is set to the value of the R component of the image.
- `VK_COMPONENT_SWIZZLE_G` specifies that the component is set to the value of the G component of the image.
- `VK_COMPONENT_SWIZZLE_B` specifies that the component is set to the value of the B component of the image.
- `VK_COMPONENT_SWIZZLE_A` specifies that the component is set to the value of the A component of the image.

Setting the identity swizzle on a component is equivalent to setting the identity mapping on that component. That is:

Table 1. Component Mappings Equivalent To `VK_COMPONENT_SWIZZLE_IDENTITY`   Component Identity Mapping

`components.r`

`VK_COMPONENT_SWIZZLE_R`

`components.g`

`VK_COMPONENT_SWIZZLE_G`

`components.b`

`VK_COMPONENT_SWIZZLE_B`

`components.a`

`VK_COMPONENT_SWIZZLE_A`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkComponentSwizzle).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0