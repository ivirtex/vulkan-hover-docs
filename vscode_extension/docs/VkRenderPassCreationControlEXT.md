# VkRenderPassCreationControlEXT(3) Manual Page

## Name

VkRenderPassCreationControlEXT - Control about the creation of render pass or subpass



## [](#_c_specification)C Specification

A `VkRenderPassCreationControlEXT` structure **can** be included in the `pNext` chain of `VkRenderPassCreateInfo2` or `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html). The `VkRenderPassCreationControlEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationControlEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           disallowMerging;
} VkRenderPassCreationControlEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `disallowMerging` is a boolean value indicating whether subpass merging will be disabled.

## [](#_description)Description

If a `VkRenderPassCreationControlEXT` structure is included in the `pNext` chain of [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html) and its value of `disallowMerging` is `VK_TRUE`, the implementation will disable subpass merging for the entire render pass. If a `VkRenderPassCreationControlEXT` structure is included in the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) and its value of `disallowMerging` is `VK_TRUE`, the implementation will disable merging the described subpass with previous subpasses in the render pass.

Valid Usage (Implicit)

- [](#VUID-VkRenderPassCreationControlEXT-sType-sType)VUID-VkRenderPassCreationControlEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_CONTROL_EXT`

## [](#_see_also)See Also

[VK\_EXT\_subpass\_merge\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subpass_merge_feedback.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), [vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassCreationControlEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0