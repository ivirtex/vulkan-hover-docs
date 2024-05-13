# VkSubpassResolvePerformanceQueryEXT(3) Manual Page

## Name

VkSubpassResolvePerformanceQueryEXT - Structure specifying the
efficiency of subpass resolve for an attachment with a format



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the performance characteristics of a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-subpass"
target="_blank" rel="noopener">subpass resolve</a> operation for an
attachment with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), add a
[VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassResolvePerformanceQueryEXT.html)
structure to the `pNext` chain of
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html).

The
[VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassResolvePerformanceQueryEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkSubpassResolvePerformanceQueryEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           optimal;
} VkSubpassResolvePerformanceQueryEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `optimal` specifies that a subpass resolve operation is optimally
  performed.

## <a href="#_description" class="anchor"></a>Description

If `optimal` is `VK_FALSE` for a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), using a
subpass resolve operation on a multisampled attachment with this format
can incur additional costs, including additional memory bandwidth usage
and a higher memory footprint. If an attachment with such a format is
used in a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
subpass, the additional memory and memory bandwidth usage can nullify
the benefits of using the
[`VK_EXT_multisampled_render_to_single_sampled`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multisampled_render_to_single_sampled.html)
extension.

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassResolvePerformanceQueryEXT-sType-sType"
  id="VUID-VkSubpassResolvePerformanceQueryEXT-sType-sType"></a>
  VUID-VkSubpassResolvePerformanceQueryEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SUBPASS_RESOLVE_PERFORMANCE_QUERY_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multisampled_render_to_single_sampled](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multisampled_render_to_single_sampled.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassResolvePerformanceQueryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
