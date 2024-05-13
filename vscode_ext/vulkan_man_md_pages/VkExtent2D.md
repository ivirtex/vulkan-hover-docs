# VkExtent2D(3) Manual Page

## Name

VkExtent2D - Structure specifying a two-dimensional extent



## <a href="#_c_specification" class="anchor"></a>C Specification

A two-dimensional extent is defined by the structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkExtent2D {
    uint32_t    width;
    uint32_t    height;
} VkExtent2D;
```

## <a href="#_members" class="anchor"></a>Members

- `width` is the width of the extent.

- `height` is the height of the extent.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html),
[VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilitiesKHR.html),
[VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPropertiesKHR.html),
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html),
[VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html),
[VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html),
[VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html),
[VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM.html),
[VkPhysicalDeviceFragmentDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapPropertiesEXT.html),
[VkPhysicalDeviceFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateKHR.html),
[VkPhysicalDeviceFragmentShadingRatePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRatePropertiesKHR.html),
[VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html),
[VkPhysicalDeviceImageProcessingPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessingPropertiesQCOM.html),
[VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html),
[VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html),
[VkPhysicalDeviceShadingRateImagePropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShadingRateImagePropertiesNV.html),
[VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html),
[VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html),
[VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html),
[VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html),
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html),
[VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
[VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html),
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html),
[VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html),
[VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html),
[VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html),
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html),
[vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html),
[vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI.html),
[vkGetRenderAreaGranularity](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html),
[vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderingAreaGranularityKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExtent2D"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
