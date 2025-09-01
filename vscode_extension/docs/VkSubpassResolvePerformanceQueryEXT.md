# VkSubpassResolvePerformanceQueryEXT(3) Manual Page

## Name

VkSubpassResolvePerformanceQueryEXT - Structure specifying the efficiency of subpass resolve for an attachment with a format



## [](#_c_specification)C Specification

To query the performance characteristics of a [subpass resolve](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-subpass) operation for an attachment with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), add a [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassResolvePerformanceQueryEXT.html) structure to the `pNext` chain of [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html).

The [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassResolvePerformanceQueryEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkSubpassResolvePerformanceQueryEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           optimal;
} VkSubpassResolvePerformanceQueryEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `optimal` specifies that a subpass resolve operation is optimally performed.

## [](#_description)Description

If `optimal` is `VK_FALSE` for a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), using a subpass resolve operation on a multisampled attachment with this format can incur additional costs, including additional memory bandwidth usage and a higher memory footprint. If an attachment with such a format is used in a [multisampled-render-to-single-sampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#subpass-multisampledrendertosinglesampled) subpass, the additional memory and memory bandwidth usage can nullify the benefits of using the `VK_EXT_multisampled_render_to_single_sampled` extension.

Valid Usage (Implicit)

- [](#VUID-VkSubpassResolvePerformanceQueryEXT-sType-sType)VUID-VkSubpassResolvePerformanceQueryEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_RESOLVE_PERFORMANCE_QUERY_EXT`

## [](#_see_also)See Also

[VK\_EXT\_multisampled\_render\_to\_single\_sampled](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multisampled_render_to_single_sampled.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassResolvePerformanceQueryEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0