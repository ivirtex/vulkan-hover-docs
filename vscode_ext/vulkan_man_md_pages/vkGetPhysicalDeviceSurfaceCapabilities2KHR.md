# vkGetPhysicalDeviceSurfaceCapabilities2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceCapabilities2KHR - Reports capabilities of a
surface on a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the basic capabilities of a surface defined by the core or
extensions, call:

``` c
// Provided by VK_KHR_get_surface_capabilities2
VkResult vkGetPhysicalDeviceSurfaceCapabilities2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    VkSurfaceCapabilities2KHR*                  pSurfaceCapabilities);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pSurfaceInfo` is a pointer to a
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure describing the surface and other fixed parameters that would
  be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pSurfaceCapabilities` is a pointer to a
  [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html) structure
  in which the capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceSurfaceCapabilities2KHR` behaves similarly to
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html),
with the ability to specify extended inputs via chained input
structures, and to return extended information via chained output
structures.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06521"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06521"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06521  
  If the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `pSurfaceInfo->surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06522"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06522"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface`
  **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

<!-- -->

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-02671"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-02671"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-02671  
  If a
  [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html)
  structure is included in the `pNext` chain of `pSurfaceCapabilities`,
  a
  [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)
  structure **must** be included in the `pNext` chain of `pSurfaceInfo`

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07776"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07776"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07776  
  If a
  [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html)
  structure is included in the `pNext` chain of `pSurfaceCapabilities`,
  a [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) structure
  **must** be included in the `pNext` chain of `pSurfaceInfo`

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07777"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07777"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07777  
  If a
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)
  structure is included in the `pNext` chain of `pSurfaceCapabilities`,
  a [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) structure
  **must** be included in the `pNext` chain of `pSurfaceInfo`

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07778"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07778"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07778  
  If a
  [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html)
  structure is included in the `pNext` chain of `pSurfaceCapabilities`,
  `pSurfaceInfo->surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07779"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07779"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07779  
  If a
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)
  structure is included in the `pNext` chain of `pSurfaceCapabilities`,
  `pSurfaceInfo->surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceCapabilities-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceCapabilities-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceCapabilities-parameter  
  `pSurfaceCapabilities` **must** be a valid pointer to a
  [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfaceCapabilities2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
