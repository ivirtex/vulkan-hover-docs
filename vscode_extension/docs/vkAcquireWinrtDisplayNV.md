# vkAcquireWinrtDisplayNV(3) Manual Page

## Name

vkAcquireWinrtDisplayNV - Acquire access to a VkDisplayKHR



## [](#_c_specification)C Specification

To acquire permission to directly access a display in Vulkan on Windows 10, call:

```c++
// Provided by VK_NV_acquire_winrt_display
VkResult vkAcquireWinrtDisplayNV(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display);
```

## [](#_parameters)Parameters

- `physicalDevice` The physical device the display is on.
- `display` The display the caller wishes to control in Vulkan.

## [](#_description)Description

All permissions necessary to control the display are granted to the Vulkan instance associated with `physicalDevice` until the display is released or the application is terminated. Permission to access the display **may** be revoked by events that cause Windows 10 itself to lose access to `display`. If this has happened, operations which require access to the display **must** fail with an appropriate error code. If permission to access `display` has already been acquired by another entity, the call **must** return the error code `VK_ERROR_INITIALIZATION_FAILED`.

Note

The Vulkan instance acquires control of a [“winrt::Windows::Devices::Display::Core::DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget) by performing an operation equivalent to [“winrt::Windows::Devices::Display::Core::DisplayManager.TryAcquireTarget()”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget) on the “DisplayTarget”.

Note

One example of when Windows 10 loses access to a display is when the display is hot-unplugged.

Note

One example of when a display has already been acquired by another entity is when the Windows desktop compositor (DWM) is in control of the display. Beginning with Windows 10 version 2004 it is possible to cause DWM to release a display by using the “Advanced display settings” sub-page of the “Display settings” control panel. [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireWinrtDisplayNV.html) does not itself cause DWM to release a display; this action must be performed outside of Vulkan.

Valid Usage (Implicit)

- [](#VUID-vkAcquireWinrtDisplayNV-physicalDevice-parameter)VUID-vkAcquireWinrtDisplayNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkAcquireWinrtDisplayNV-display-parameter)VUID-vkAcquireWinrtDisplayNV-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkAcquireWinrtDisplayNV-display-parent)VUID-vkAcquireWinrtDisplayNV-display-parent  
  `display` **must** have been created, allocated, or retrieved from `physicalDevice`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_acquire\_winrt\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_acquire_winrt_display.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireWinrtDisplayNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0