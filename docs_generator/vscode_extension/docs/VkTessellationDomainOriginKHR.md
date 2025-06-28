# VkTessellationDomainOrigin(3) Manual Page

## Name

VkTessellationDomainOrigin - Enum describing tessellation domain origin



## [](#_c_specification)C Specification

The possible tessellation domain origins are specified by the [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html) enumeration:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkTessellationDomainOrigin {
    VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT = 0,
    VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT = 1,
  // Provided by VK_KHR_maintenance2
    VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR = VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT,
  // Provided by VK_KHR_maintenance2
    VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR = VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT,
} VkTessellationDomainOrigin;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance2
typedef VkTessellationDomainOrigin VkTessellationDomainOriginKHR;
```

## [](#_description)Description

- `VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT` specifies that the origin of the domain space is in the upper left corner, as shown in figure [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#img-tessellation-topology-ul](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#img-tessellation-topology-ul).
- `VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT` specifies that the origin of the domain space is in the lower left corner, as shown in figure [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#img-tessellation-topology-ll](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#img-tessellation-topology-ll).

This enum affects how the `VertexOrderCw` and `VertexOrderCcw` tessellation execution modes are interpreted, since the winding is defined relative to the orientation of the domain.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html), [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetTessellationDomainOriginEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTessellationDomainOrigin)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0