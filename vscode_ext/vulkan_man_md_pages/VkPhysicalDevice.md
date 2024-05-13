# VkPhysicalDevice(3) Manual Page

## Name

VkPhysicalDevice - Opaque handle to a physical device object



## <a href="#_c_specification" class="anchor"></a>C Specification

Vulkan separates the concept of *physical* and *logical* devices. A
physical device usually represents a single complete implementation of
Vulkan (excluding instance-level functionality) available to the host,
of which there are a finite number. A logical device represents an
instance of that implementation with its own state and resources
independent of other logical devices.

Physical devices are represented by `VkPhysicalDevice` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_HANDLE(VkPhysicalDevice)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html),
[VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html),
[vkAcquireDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireDrmDisplayEXT.html),
[vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html),
[vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireXlibDisplayEXT.html),
[vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html),
[vkCreateDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayModeKHR.html),
[vkEnumerateDeviceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html),
[vkEnumerateDeviceLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceLayerProperties.html),
[vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html),
[vkEnumeratePhysicalDevices](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDevices.html),
[vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModeProperties2KHR.html),
[vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModePropertiesKHR.html),
[vkGetDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilities2KHR.html),
[vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html),
[vkGetDisplayPlaneSupportedDisplaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html),
[vkGetDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDrmDisplayEXT.html),
[vkGetPhysicalDeviceCalibrateableTimeDomainsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsEXT.html),
[vkGetPhysicalDeviceCalibrateableTimeDomainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCalibrateableTimeDomainsKHR.html),
[vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.html),
[vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html),
[vkGetPhysicalDeviceDirectFBPresentationSupportEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDirectFBPresentationSupportEXT.html),
[vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html),
[vkGetPhysicalDeviceDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html),
[vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html),
[vkGetPhysicalDeviceDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPropertiesKHR.html),
[vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html),
[vkGetPhysicalDeviceExternalBufferPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferPropertiesKHR.html),
[vkGetPhysicalDeviceExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html),
[vkGetPhysicalDeviceExternalFencePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFencePropertiesKHR.html),
[vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html),
[vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html),
[vkGetPhysicalDeviceExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html),
[vkGetPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures.html),
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html),
[vkGetPhysicalDeviceFeatures2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2KHR.html),
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html),
[vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html),
[vkGetPhysicalDeviceFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html),
[vkGetPhysicalDeviceFragmentShadingRatesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFragmentShadingRatesKHR.html),
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html),
[vkGetPhysicalDeviceImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html),
[vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html),
[vkGetPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html),
[vkGetPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2KHR.html),
[vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html),
[vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html),
[vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html),
[vkGetPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html),
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
[vkGetPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2KHR.html),
[vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html),
[vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html),
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html),
[vkGetPhysicalDeviceQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2KHR.html),
[vkGetPhysicalDeviceScreenPresentationSupportQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceScreenPresentationSupportQNX.html),
[vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html),
[vkGetPhysicalDeviceSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2.html),
[vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html),
[vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html),
[vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html),
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html),
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html),
[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html),
[vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html),
[vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html),
[vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html),
[vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html),
[vkGetPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceToolProperties.html),
[vkGetPhysicalDeviceToolPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceToolPropertiesEXT.html),
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html),
[vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html),
[vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html),
[vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html),
[vkGetPhysicalDeviceWin32PresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html),
[vkGetPhysicalDeviceXcbPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceXcbPresentationSupportKHR.html),
[vkGetPhysicalDeviceXlibPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceXlibPresentationSupportKHR.html),
[vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRandROutputDisplayEXT.html),
[vkGetWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetWinrtDisplayNV.html),
[vkReleaseDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseDisplayEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevice"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
