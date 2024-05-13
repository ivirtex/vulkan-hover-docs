# VkChromaLocation(3) Manual Page

## Name

VkChromaLocation - Position of downsampled chroma samples



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html) enum defines the location
of downsampled chroma component samples relative to the luma samples,
and is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkChromaLocation {
    VK_CHROMA_LOCATION_COSITED_EVEN = 0,
    VK_CHROMA_LOCATION_MIDPOINT = 1,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_CHROMA_LOCATION_COSITED_EVEN_KHR = VK_CHROMA_LOCATION_COSITED_EVEN,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_CHROMA_LOCATION_MIDPOINT_KHR = VK_CHROMA_LOCATION_MIDPOINT,
} VkChromaLocation;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkChromaLocation VkChromaLocationKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_CHROMA_LOCATION_COSITED_EVEN` specifies that downsampled chroma
  samples are aligned with luma samples with even coordinates.

- `VK_CHROMA_LOCATION_MIDPOINT` specifies that downsampled chroma
  samples are located half way between each even luma sample and the
  nearest higher odd luma sample.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html),
[VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html),
[VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkChromaLocation"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
