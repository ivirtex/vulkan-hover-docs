# VkAttachmentFeedbackLoopInfoEXT(3) Manual Page

## Name

VkAttachmentFeedbackLoopInfoEXT - Structure specifying whether feedback loop is enabled for an attachment



## [](#_c_specification)C Specification

To [enable feedback loop](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-feedbackloop) for an attachment, the `VkAttachmentFeedbackLoopInfoEXT` structure **can** be added to the `pNext` chain of [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html).

The `VkAttachmentFeedbackLoopInfoEXT` structure is defined as:

```c++
// Provided by VK_KHR_unified_image_layouts with VK_EXT_attachment_feedback_loop_layout and (VK_VERSION_1_3 or VK_KHR_dynamic_rendering)
typedef struct VkAttachmentFeedbackLoopInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           feedbackLoopEnable;
} VkAttachmentFeedbackLoopInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `feedbackLoopEnable` specifies that [feedback loop is enabled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-feedbackloop) for the attachment identified by [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html)::`imageView`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAttachmentFeedbackLoopInfoEXT-unifiedImageLayouts-10782)VUID-VkAttachmentFeedbackLoopInfoEXT-unifiedImageLayouts-10782  
  If the [`unifiedImageLayouts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-unifiedImageLayouts) feature is not enabled, `feedbackLoopEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentFeedbackLoopInfoEXT-sType-sType)VUID-VkAttachmentFeedbackLoopInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_FEEDBACK_LOOP_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_attachment\_feedback\_loop\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_layout.html), [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_KHR\_unified\_image\_layouts](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_unified_image_layouts.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentFeedbackLoopInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0