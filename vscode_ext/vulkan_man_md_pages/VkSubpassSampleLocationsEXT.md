# VkSubpassSampleLocationsEXT(3) Manual Page

## Name

VkSubpassSampleLocationsEXT - Structure specifying the sample locations
state to use for layout transitions of attachments performed after a
given subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassSampleLocationsEXT` structure is defined as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkSubpassSampleLocationsEXT {
    uint32_t                    subpassIndex;
    VkSampleLocationsInfoEXT    sampleLocationsInfo;
} VkSubpassSampleLocationsEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `subpassIndex` is the index of the subpass for which the sample
  locations state is provided.

- `sampleLocationsInfo` is the sample locations state to use for the
  layout transition of the depth/stencil attachment away from the image
  layout the attachment is used with in the subpass specified in
  `subpassIndex`.

## <a href="#_description" class="anchor"></a>Description

If the image referenced by the depth/stencil attachment used in the
subpass identified by `subpassIndex` was not created with
`VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT` or if the
subpass does not use a depth/stencil attachment, and
[VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations`
is `VK_TRUE` then the values specified in `sampleLocationsInfo` are
ignored.

Valid Usage

- <a href="#VUID-VkSubpassSampleLocationsEXT-subpassIndex-01532"
  id="VUID-VkSubpassSampleLocationsEXT-subpassIndex-01532"></a>
  VUID-VkSubpassSampleLocationsEXT-subpassIndex-01532  
  `subpassIndex` **must** be less than the `subpassCount` specified in
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) the render pass
  specified by
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderPass` was
  created with

Valid Usage (Implicit)

- <a
  href="#VUID-VkSubpassSampleLocationsEXT-sampleLocationsInfo-parameter"
  id="VUID-VkSubpassSampleLocationsEXT-sampleLocationsInfo-parameter"></a>
  VUID-VkSubpassSampleLocationsEXT-sampleLocationsInfo-parameter  
  `sampleLocationsInfo` **must** be a valid
  [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassSampleLocationsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
