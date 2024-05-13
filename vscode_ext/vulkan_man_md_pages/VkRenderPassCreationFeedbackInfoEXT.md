# VkRenderPassCreationFeedbackInfoEXT(3) Manual Page

## Name

VkRenderPassCreationFeedbackInfoEXT - Feedback about the creation of a
render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassCreationFeedbackInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationFeedbackInfoEXT {
    uint32_t    postMergeSubpassCount;
} VkRenderPassCreationFeedbackInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `postMergeSubpassCount` is the subpass count after merge.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkRenderPassCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreationFeedbackInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
