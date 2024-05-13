# VkCoarseSampleOrderCustomNV(3) Manual Page

## Name

VkCoarseSampleOrderCustomNV - Structure specifying parameters
controlling shading rate image usage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCoarseSampleOrderCustomNV` structure is defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkCoarseSampleOrderCustomNV {
    VkShadingRatePaletteEntryNV        shadingRate;
    uint32_t                           sampleCount;
    uint32_t                           sampleLocationCount;
    const VkCoarseSampleLocationNV*    pSampleLocations;
} VkCoarseSampleOrderCustomNV;
```

## <a href="#_members" class="anchor"></a>Members

- `shadingRate` is a shading rate palette entry that identifies the
  fragment width and height for the combination of fragment area and
  per-pixel coverage sample count to control.

- `sampleCount` identifies the per-pixel coverage sample count for the
  combination of fragment area and coverage sample count to control.

- `sampleLocationCount` specifies the number of sample locations in the
  custom ordering.

- `pSampleLocations` is a pointer to an array of
  [VkCoarseSampleLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleLocationNV.html) structures
  specifying the location of each sample in the custom ordering.

## <a href="#_description" class="anchor"></a>Description

The `VkCoarseSampleOrderCustomNV` structure is used with a coverage
sample ordering type of `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV` to
specify the order of coverage samples for one combination of fragment
width, fragment height, and coverage sample count.

When using a custom sample ordering, element *j* in `pSampleLocations`
specifies a specific pixel location and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">sample index</a> that corresponds to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">coverage index</a> *j* in the multi-pixel
fragment.

Valid Usage

- <a href="#VUID-VkCoarseSampleOrderCustomNV-shadingRate-02073"
  id="VUID-VkCoarseSampleOrderCustomNV-shadingRate-02073"></a>
  VUID-VkCoarseSampleOrderCustomNV-shadingRate-02073  
  `shadingRate` **must** be a shading rate that generates fragments with
  more than one pixel

- <a href="#VUID-VkCoarseSampleOrderCustomNV-sampleCount-02074"
  id="VUID-VkCoarseSampleOrderCustomNV-sampleCount-02074"></a>
  VUID-VkCoarseSampleOrderCustomNV-sampleCount-02074  
  `sampleCount` **must** correspond to a sample count enumerated in
  [VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html) whose corresponding bit
  is set in
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`framebufferNoAttachmentsSampleCounts`

- <a href="#VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02075"
  id="VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02075"></a>
  VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02075  
  `sampleLocationCount` **must** be equal to the product of
  `sampleCount`, the fragment width for `shadingRate`, and the fragment
  height for `shadingRate`

- <a href="#VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02076"
  id="VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02076"></a>
  VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-02076  
  `sampleLocationCount` **must** be less than or equal to the value of
  `VkPhysicalDeviceShadingRateImagePropertiesNV`::`shadingRateMaxCoarseSamples`

- <a href="#VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-02077"
  id="VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-02077"></a>
  VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-02077  
  The array `pSampleLocations` **must** contain exactly one entry for
  every combination of valid values for `pixelX`, `pixelY`, and `sample`
  in the structure
  [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)

Valid Usage (Implicit)

- <a href="#VUID-VkCoarseSampleOrderCustomNV-shadingRate-parameter"
  id="VUID-VkCoarseSampleOrderCustomNV-shadingRate-parameter"></a>
  VUID-VkCoarseSampleOrderCustomNV-shadingRate-parameter  
  `shadingRate` **must** be a valid
  [VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteEntryNV.html) value

- <a href="#VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-parameter"
  id="VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-parameter"></a>
  VUID-VkCoarseSampleOrderCustomNV-pSampleLocations-parameter  
  `pSampleLocations` **must** be a valid pointer to an array of
  `sampleLocationCount`
  [VkCoarseSampleLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleLocationNV.html) structures

- <a
  href="#VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-arraylength"
  id="VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-arraylength"></a>
  VUID-VkCoarseSampleOrderCustomNV-sampleLocationCount-arraylength  
  `sampleLocationCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkCoarseSampleLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleLocationNV.html),
[VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html),
[VkShadingRatePaletteEntryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteEntryNV.html),
[vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCoarseSampleOrderCustomNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
