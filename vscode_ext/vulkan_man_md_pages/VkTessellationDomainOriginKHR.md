# VkTessellationDomainOrigin(3) Manual Page

## Name

VkTessellationDomainOrigin - Enum describing tessellation domain origin



## <a href="#_c_specification" class="anchor"></a>C Specification

The possible tessellation domain origins are specified by the
[VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html)
enumeration:

``` c
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

``` c
// Provided by VK_KHR_maintenance2
typedef VkTessellationDomainOrigin VkTessellationDomainOriginKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT` specifies that the origin
  of the domain space is in the upper left corner, as shown in figure <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#img-tessellation-topology-ul"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#img-tessellation-topology-ul</a>.

- `VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT` specifies that the origin
  of the domain space is in the lower left corner, as shown in figure <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#img-tessellation-topology-ll"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#img-tessellation-topology-ll</a>.

This enum affects how the `VertexOrderCw` and `VertexOrderCcw`
tessellation execution modes are interpreted, since the winding is
defined relative to the orientation of the domain.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html),
[vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTessellationDomainOrigin"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
