# VkAttachmentSampleLocationsEXT(3) Manual Page

## Name

VkAttachmentSampleLocationsEXT - Structure specifying the sample
locations state to use in the initial layout transition of attachments



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentSampleLocationsEXT` structure is defined as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkAttachmentSampleLocationsEXT {
    uint32_t                    attachmentIndex;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkAttachmentSampleLocationsEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `attachmentIndex` is the index of the attachment for which the sample
  locations state is provided.

- `sampleLocationsInfo` is the sample locations state to use for the
  layout transition of the given attachment from the initial layout of
  the attachment to the image layout specified for the attachment in the
  first subpass using it.

## <a href="#_description" class="anchor"></a>Description

If the image referenced by the framebuffer attachment at index
`attachmentIndex` was not created with
`VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` then the
values specified in `sampleLocationsInfo` are ignored.

Valid Usage

- <a href="#VUID-VkAttachmentSampleLocationsEXT-attachmentIndex-01531"
  id="VUID-VkAttachmentSampleLocationsEXT-attachmentIndex-01531"></a>
  VUID-VkAttachmentSampleLocationsEXT-attachmentIndex-01531  
  `attachmentIndex` **must** be less than the `attachmentCount`
  specified in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) the
  render pass specified by
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderPass` was
  created with

Valid Usage (Implicit)

- <a
  href="#VUID-VkAttachmentSampleLocationsEXT-sampleLocationsInfo-parameter"
  id="VUID-VkAttachmentSampleLocationsEXT-sampleLocationsInfo-parameter"></a>
  VUID-VkAttachmentSampleLocationsEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid
  [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentSampleLocationsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
