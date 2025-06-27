# vkGetWinrtDisplayNV(3) Manual Page

## Name

vkGetWinrtDisplayNV - Query the VkDisplayKHR corresponding to a WinRT
DisplayTarget



## <a href="#_c_specification" class="anchor"></a>C Specification

When acquiring displays on Windows 10, an application may also wish to
enumerate and identify them using a native handle rather than a
`VkDisplayKHR` handle.

To determine the `VkDisplayKHR` handle corresponding to a
[“winrt::Windows::Devices::Display::Core::DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget),
call:

``` c
// Provided by VK_NV_acquire_winrt_display
VkResult vkGetWinrtDisplayNV(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    deviceRelativeId,
    VkDisplayKHR*                               pDisplay);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device on which to query the display
  handle.

- `deviceRelativeId` The value of the
  [“AdapterRelativeId”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget.adapterrelativeid)
  property of a
  [“DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget)
  that is enumerated by a
  [“DisplayAdapter”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displayadapter)
  with an
  [“Id”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displayadapter.id)
  property matching the `deviceLUID` property of a
  [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html) for
  `physicalDevice`.

- `pDisplay` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle
  will be returned here.

## <a href="#_description" class="anchor"></a>Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) corresponding to
`deviceRelativeId` on `physicalDevice`,
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be returned in
`pDisplay`.

Valid Usage (Implicit)

- <a href="#VUID-vkGetWinrtDisplayNV-physicalDevice-parameter"
  id="VUID-vkGetWinrtDisplayNV-physicalDevice-parameter"></a>
  VUID-vkGetWinrtDisplayNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetWinrtDisplayNV-pDisplay-parameter"
  id="VUID-vkGetWinrtDisplayNV-pDisplay-parameter"></a>
  VUID-vkGetWinrtDisplayNV-pDisplay-parameter  
  `pDisplay` **must** be a valid pointer to a
  [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_acquire_winrt_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_acquire_winrt_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetWinrtDisplayNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
