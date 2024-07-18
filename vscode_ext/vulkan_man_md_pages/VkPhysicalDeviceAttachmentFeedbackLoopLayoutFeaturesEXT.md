# VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT - Structure
indicating support for a render feedback loop image layout



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT` structure
is defined as:

``` c
// Provided by VK_EXT_attachment_feedback_loop_layout
typedef struct VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           attachmentFeedbackLoopLayout;
} VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-attachmentFeedbackLoopLayout"></span>
  `attachmentFeedbackLoopLayout` indicates whether the implementation
  supports using `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
  image layout for images created with
  `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ATTACHMENT_FEEDBACK_LOOP_LAYOUT_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_attachment_feedback_loop_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_attachment_feedback_loop_layout.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
