# VkSubpassMergeStatusEXT(3) Manual Page

## Name

VkSubpassMergeStatusEXT - Specify a subpass merging status



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
`VkRenderPassSubpassFeedbackInfoEXT`:subpassMergeStatus are:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef enum VkSubpassMergeStatusEXT {
    VK_SUBPASS_MERGE_STATUS_MERGED_EXT = 0,
    VK_SUBPASS_MERGE_STATUS_DISALLOWED_EXT = 1,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SIDE_EFFECTS_EXT = 2,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SAMPLES_MISMATCH_EXT = 3,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_VIEWS_MISMATCH_EXT = 4,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_ALIASING_EXT = 5,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_DEPENDENCIES_EXT = 6,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_INCOMPATIBLE_INPUT_ATTACHMENT_EXT = 7,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_TOO_MANY_ATTACHMENTS_EXT = 8,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_INSUFFICIENT_STORAGE_EXT = 9,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_DEPTH_STENCIL_COUNT_EXT = 10,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_RESOLVE_ATTACHMENT_REUSE_EXT = 11,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SINGLE_SUBPASS_EXT = 12,
    VK_SUBPASS_MERGE_STATUS_NOT_MERGED_UNSPECIFIED_EXT = 13,
} VkSubpassMergeStatusEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SUBPASS_MERGE_STATUS_MERGED_EXT` specifies the subpass is merged
  with a previous subpass.

- `VK_SUBPASS_MERGE_STATUS_DISALLOWED_EXT` specifies the subpass is
  disallowed to merge with previous subpass. If the render pass does not
  allow subpass merging, then all subpass statuses are set to this
  value. If a subpass description does not allow subpass merging, then
  only that subpass’s status is set to this value.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SIDE_EFFECTS_EXT` specifies the
  subpass is not merged because it contains side effects.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SAMPLES_MISMATCH_EXT` specifies
  the subpass is not merged because sample count is not compatible with
  previous subpass.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_VIEWS_MISMATCH_EXT` specifies the
  subpass is not merged because view masks do not match with previous
  subpass.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_ALIASING_EXT` specifies the
  subpass is not merged because of attachments aliasing between them.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_DEPENDENCIES_EXT` specifies the
  subpass is not merged because subpass dependencies do not allow
  merging.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_INCOMPATIBLE_INPUT_ATTACHMENT_EXT`
  specifies the subpass is not merged because input attachment is not a
  color attachment from previous subpass or the formats are
  incompatible.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_TOO_MANY_ATTACHMENTS_EXT`
  specifies the subpass is not merged because of too many attachments.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_INSUFFICIENT_STORAGE_EXT`
  specifies the subpass is not merged because of insufficient memory.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_DEPTH_STENCIL_COUNT_EXT` specifies
  the subpass is not merged because of too many depth/stencil
  attachments.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_RESOLVE_ATTACHMENT_REUSE_EXT`
  specifies the subpass is not merged because a resolve attachment is
  reused as an input attachment in a subsequent subpass.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_SINGLE_SUBPASS_EXT` specifies the
  subpass is not merged because the render pass has only one subpass.

- `VK_SUBPASS_MERGE_STATUS_NOT_MERGED_UNSPECIFIED_EXT` specifies other
  reasons why subpass is not merged. It is also the recommended default
  value that should be reported when a subpass is not merged and when no
  other value is appropriate.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassMergeStatusEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
