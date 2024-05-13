# vkAcquireDrmDisplayEXT(3) Manual Page

## Name

vkAcquireDrmDisplayEXT - Acquire access to a VkDisplayKHR using DRM



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire permission to directly a display in Vulkan from the Direct
Rendering Manager (DRM) interface, call:

``` c
// Provided by VK_EXT_acquire_drm_display
VkResult vkAcquireDrmDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    int32_t                                     drmFd,
    VkDisplayKHR                                display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device the display is on.

- `drmFd` DRM primary file descriptor.

- `display` The display the caller wishes Vulkan to control.

## <a href="#_description" class="anchor"></a>Description

All permissions necessary to control the display are granted to the
Vulkan instance associated with the provided `physicalDevice` until the
display is either released or the connector is unplugged. The provided
`drmFd` must correspond to the one owned by the `physicalDevice`. If
not, the error code `VK_ERROR_UNKNOWN` must be returned. The DRM FD must
have DRM master permissions. If any error is encountered during the
acquisition of the display, the call must return the error code
`VK_ERROR_INITIALIZATION_FAILED`.

The provided DRM fd should not be closed before the display is released,
attempting to do it may result in undefined behavior.

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireDrmDisplayEXT-physicalDevice-parameter"
  id="VUID-vkAcquireDrmDisplayEXT-physicalDevice-parameter"></a>
  VUID-vkAcquireDrmDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkAcquireDrmDisplayEXT-display-parameter"
  id="VUID-vkAcquireDrmDisplayEXT-display-parameter"></a>
  VUID-vkAcquireDrmDisplayEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkAcquireDrmDisplayEXT-display-parent"
  id="VUID-vkAcquireDrmDisplayEXT-display-parent"></a>
  VUID-vkAcquireDrmDisplayEXT-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_acquire_drm_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_acquire_drm_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireDrmDisplayEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
