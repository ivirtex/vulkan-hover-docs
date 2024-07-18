# vkGetDisplayPlaneSupportedDisplaysKHR(3) Manual Page

## Name

vkGetDisplayPlaneSupportedDisplaysKHR - Query the list of displays a
plane supports



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine which displays a plane is usable with, call

``` c
// Provided by VK_KHR_display
VkResult vkGetDisplayPlaneSupportedDisplaysKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    planeIndex,
    uint32_t*                                   pDisplayCount,
    VkDisplayKHR*                               pDisplays);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is a physical device.

- `planeIndex` is the plane which the application wishes to use, and
  **must** be in the range \[0, physical device plane count - 1\].

- `pDisplayCount` is a pointer to an integer related to the number of
  displays available or queried, as described below.

- `pDisplays` is either `NULL` or a pointer to an array of
  `VkDisplayKHR` handles.

## <a href="#_description" class="anchor"></a>Description

If `pDisplays` is `NULL`, then the number of displays usable with the
specified `planeIndex` for `physicalDevice` is returned in
`pDisplayCount`. Otherwise, `pDisplayCount` **must** point to a variable
set by the application to the number of elements in the `pDisplays`
array, and on return the variable is overwritten with the number of
handles actually written to `pDisplays`. If the value of `pDisplayCount`
is less than the number of usable display-plane pairs for
`physicalDevice`, at most `pDisplayCount` handles will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available pairs were returned.

Valid Usage

- <a href="#VUID-vkGetDisplayPlaneSupportedDisplaysKHR-planeIndex-01249"
  id="VUID-vkGetDisplayPlaneSupportedDisplaysKHR-planeIndex-01249"></a>
  VUID-vkGetDisplayPlaneSupportedDisplaysKHR-planeIndex-01249  
  `planeIndex` **must** be less than the number of display planes
  supported by the device as determined by calling
  `vkGetPhysicalDeviceDisplayPlanePropertiesKHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDisplayPlaneSupportedDisplaysKHR-physicalDevice-parameter"
  id="VUID-vkGetDisplayPlaneSupportedDisplaysKHR-physicalDevice-parameter"></a>
  VUID-vkGetDisplayPlaneSupportedDisplaysKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplayCount-parameter"
  id="VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplayCount-parameter"></a>
  VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplayCount-parameter  
  `pDisplayCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplays-parameter"
  id="VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplays-parameter"></a>
  VUID-vkGetDisplayPlaneSupportedDisplaysKHR-pDisplays-parameter  
  If the value referenced by `pDisplayCount` is not `0`, and `pDisplays`
  is not `NULL`, `pDisplays` **must** be a valid pointer to an array of
  `pDisplayCount` [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handles

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDisplayPlaneSupportedDisplaysKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
