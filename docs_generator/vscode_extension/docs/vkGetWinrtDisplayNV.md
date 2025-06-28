# vkGetWinrtDisplayNV(3) Manual Page

## Name

vkGetWinrtDisplayNV - Query the VkDisplayKHR corresponding to a WinRT DisplayTarget



## [](#_c_specification)C Specification

When acquiring displays on Windows 10, an application may also wish to enumerate and identify them using a native handle rather than a `VkDisplayKHR` handle.

To determine the `VkDisplayKHR` handle corresponding to a [“winrt::Windows::Devices::Display::Core::DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget), call:

```c++
// Provided by VK_NV_acquire_winrt_display
VkResult vkGetWinrtDisplayNV(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    deviceRelativeId,
    VkDisplayKHR*                               pDisplay);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device on which to query the display handle.
- `deviceRelativeId` The value of the [“AdapterRelativeId”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget.adapterrelativeid) property of a [“DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget) that is enumerated by a [“DisplayAdapter”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displayadapter) with an [“Id”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displayadapter.id) property matching the `deviceLUID` property of a [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html) for `physicalDevice`.
- `pDisplay` The corresponding [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle will be returned here.

## [](#_description)Description

If there is no [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) corresponding to `deviceRelativeId` on `physicalDevice`, [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** be returned in `pDisplay`.

Valid Usage (Implicit)

- [](#VUID-vkGetWinrtDisplayNV-physicalDevice-parameter)VUID-vkGetWinrtDisplayNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetWinrtDisplayNV-pDisplay-parameter)VUID-vkGetWinrtDisplayNV-pDisplay-parameter  
  `pDisplay` **must** be a valid pointer to a [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_INITIALIZATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_acquire\_winrt\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_acquire_winrt_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetWinrtDisplayNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0