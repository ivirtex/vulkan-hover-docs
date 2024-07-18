# VkRenderPassCreationControlEXT(3) Manual Page

## Name

VkRenderPassCreationControlEXT - Control about the creation of render
pass or subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

A `VkRenderPassCreationControlEXT` structure **can** be included in the
`pNext` chain of `VkRenderPassCreateInfo2` or `pNext` chain of
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html). The
`VkRenderPassCreationControlEXT` structure is defined as:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationControlEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           disallowMerging;
} VkRenderPassCreationControlEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `disallowMerging` is a boolean value indicating whether subpass
  merging will be disabled.

## <a href="#_description" class="anchor"></a>Description

If a `VkRenderPassCreationControlEXT` structure is included in the
`pNext` chain of [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html)
and its value of `disallowMerging` is `VK_TRUE`, the implementation will
disable subpass merging for the entire render pass. If a
`VkRenderPassCreationControlEXT` structure is included in the `pNext`
chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) and its
value of `disallowMerging` is `VK_TRUE`, the implementation will disable
merging the described subpass with previous subpasses in the render
pass.

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassCreationControlEXT-sType-sType"
  id="VUID-VkRenderPassCreationControlEXT-sType-sType"></a>
  VUID-VkRenderPassCreationControlEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_CONTROL_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreationControlEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
