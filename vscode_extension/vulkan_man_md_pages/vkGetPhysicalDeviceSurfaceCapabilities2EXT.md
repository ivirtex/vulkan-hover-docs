# vkGetPhysicalDeviceSurfaceCapabilities2EXT(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceCapabilities2EXT - Query surface capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the basic capabilities of a surface, needed in order to create
a swapchain, call:

``` c
// Provided by VK_EXT_display_surface_counter
VkResult vkGetPhysicalDeviceSurfaceCapabilities2EXT(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    VkSurfaceCapabilities2EXT*                  pSurfaceCapabilities);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `surface` is the surface that will be associated with the swapchain.

- `pSurfaceCapabilities` is a pointer to a
  [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html) structure
  in which the capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceSurfaceCapabilities2EXT` behaves similarly to
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html),
with the ability to return extended information by adding extending
structures to the `pNext` chain of its `pSurfaceCapabilities` parameter.

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06523"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06523"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06523  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06211"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06211"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06211  
  `surface` **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-pSurfaceCapabilities-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-pSurfaceCapabilities-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-pSurfaceCapabilities-parameter  
  `pSurfaceCapabilities` **must** be a valid pointer to a
  [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html) structure

- <a href="#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-commonparent"
  id="VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-commonparent"></a>
  VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-commonparent  
  Both of `physicalDevice`, and `surface` **must** have been created,
  allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_surface_counter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_surface_counter.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfaceCapabilities2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
