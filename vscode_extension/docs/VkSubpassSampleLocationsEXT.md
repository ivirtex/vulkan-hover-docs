# VkSubpassSampleLocationsEXT(3) Manual Page

## Name

VkSubpassSampleLocationsEXT - Structure specifying the sample locations state to use for layout transitions of attachments performed after a given subpass



## [](#_c_specification)C Specification

The `VkSubpassSampleLocationsEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkSubpassSampleLocationsEXT {
    uint32_t                    subpassIndex;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkSubpassSampleLocationsEXT;
```

## [](#_members)Members

- `subpassIndex` is the index of the subpass for which the sample locations state is provided.
- `sampleLocationsInfo` is the sample locations state to use for the layout transition of the depth/stencil attachment away from the image layout the attachment is used with in the subpass specified in `subpassIndex`.

## [](#_description)Description

If the image referenced by the depth/stencil attachment used in the subpass identified by `subpassIndex` was not created with `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` or if the subpass does not use a depth/stencil attachment, and [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations` is `VK_TRUE` then the values specified in `sampleLocationsInfo` are ignored.

Valid Usage

- [](#VUID-VkSubpassSampleLocationsEXT-subpassIndex-01532)VUID-VkSubpassSampleLocationsEXT-subpassIndex-01532  
  `subpassIndex` **must** be less than the `subpassCount` specified in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) the render pass specified by [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html)::`renderPass` was created with

Valid Usage (Implicit)

- [](#VUID-VkSubpassSampleLocationsEXT-sampleLocationsInfo-parameter)VUID-VkSubpassSampleLocationsEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html), [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassSampleLocationsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0