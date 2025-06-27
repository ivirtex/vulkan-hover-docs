# vkGetDrmDisplayEXT(3) Manual Page

## Name

vkGetDrmDisplayEXT - Query the VkDisplayKHR corresponding to a DRM
connector ID



## <a href="#_c_specification" class="anchor"></a>C Specification

Before acquiring a display from the DRM interface, the caller may want
to select a specific `VkDisplayKHR` handle by identifying it using a
`connectorId`. To do so, call:

``` c
// Provided by VK_EXT_acquire_drm_display
VkResult vkGetDrmDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    int32_t                                     drmFd,
    uint32_t                                    connectorId,
    VkDisplayKHR*                               display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device to query the display from.

- `drmFd` DRM primary file descriptor.

- `connectorId` Identifier of the specified DRM connector.

- `display` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle
  will be returned here.

## <a href="#_description" class="anchor"></a>Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) corresponding to the
`connectorId` on the `physicalDevice`, the returning `display` must be
set to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). The provided `drmFd` must
correspond to the one owned by the `physicalDevice`. If not, the error
code `VK_ERROR_UNKNOWN` must be returned. Master permissions are not
required, because the file descriptor is just used for information
gathering purposes. The given `connectorId` must be a resource owned by
the provided `drmFd`. If not, the error code `VK_ERROR_UNKNOWN` must be
returned. If any error is encountered during the identification of the
display, the call must return the error code
`VK_ERROR_INITIALIZATION_FAILED`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetDrmDisplayEXT-physicalDevice-parameter"
  id="VUID-vkGetDrmDisplayEXT-physicalDevice-parameter"></a>
  VUID-vkGetDrmDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetDrmDisplayEXT-display-parameter"
  id="VUID-vkGetDrmDisplayEXT-display-parameter"></a>
  VUID-vkGetDrmDisplayEXT-display-parameter  
  `display` **must** be a valid pointer to a
  [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_acquire_drm_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_acquire_drm_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDrmDisplayEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
