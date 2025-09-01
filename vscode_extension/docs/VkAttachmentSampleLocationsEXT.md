# VkAttachmentSampleLocationsEXT(3) Manual Page

## Name

VkAttachmentSampleLocationsEXT - Structure specifying the sample locations state to use in the initial layout transition of attachments



## [](#_c_specification)C Specification

The `VkAttachmentSampleLocationsEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkAttachmentSampleLocationsEXT {
    uint32_t                    attachmentIndex;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkAttachmentSampleLocationsEXT;
```

## [](#_members)Members

- `attachmentIndex` is the index of the attachment for which the sample locations state is provided.
- `sampleLocationsInfo` is the sample locations state to use for the layout transition of the given attachment from the initial layout of the attachment to the image layout specified for the attachment in the first subpass using it.

## [](#_description)Description

If the image referenced by the framebuffer attachment at index `attachmentIndex` was not created with `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` then the values specified in `sampleLocationsInfo` are ignored.

Valid Usage

- [](#VUID-VkAttachmentSampleLocationsEXT-attachmentIndex-01531)VUID-VkAttachmentSampleLocationsEXT-attachmentIndex-01531  
  `attachmentIndex` **must** be less than the `attachmentCount` specified in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) the render pass specified by [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html)::`renderPass` was created with

Valid Usage (Implicit)

- [](#VUID-VkAttachmentSampleLocationsEXT-sampleLocationsInfo-parameter)VUID-VkAttachmentSampleLocationsEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html), [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentSampleLocationsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0