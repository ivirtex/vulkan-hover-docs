# VkRenderPassSubpassFeedbackInfoEXT(3) Manual Page

## Name

VkRenderPassSubpassFeedbackInfoEXT - Feedback about the creation of
subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassSubpassFeedbackInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassSubpassFeedbackInfoEXT {
    VkSubpassMergeStatusEXT    subpassMergeStatus;
    char                       description[VK_MAX_DESCRIPTION_SIZE];
    uint32_t                   postMergeIndex;
} VkRenderPassSubpassFeedbackInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `subpassMergeStatus` is a `VkSubpassMergeStatusEXT` value specifying
  information about whether the subpass is merged with previous subpass
  and the reason why it is not merged.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which provides additional
  details.

- `postMergeIndex` is the subpass index after the subpass merging.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html),
[VkSubpassMergeStatusEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassMergeStatusEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassSubpassFeedbackInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
