# vkDisplayPowerControlEXT(3) Manual Page

## Name

vkDisplayPowerControlEXT - Set the power state of a display



## [](#_c_specification)C Specification

To set the power state of a display, call:

```c++
// Provided by VK_EXT_display_control
VkResult vkDisplayPowerControlEXT(
    VkDevice                                    device,
    VkDisplayKHR                                display,
    const VkDisplayPowerInfoEXT*                pDisplayPowerInfo);
```

## [](#_parameters)Parameters

- `device` is a logical device associated with `display`.
- `display` is the display whose power state is modified.
- `pDisplayPowerInfo` is a pointer to a [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerInfoEXT.html) structure specifying the new power state of `display`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDisplayPowerControlEXT-device-parameter)VUID-vkDisplayPowerControlEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDisplayPowerControlEXT-display-parameter)VUID-vkDisplayPowerControlEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkDisplayPowerControlEXT-pDisplayPowerInfo-parameter)VUID-vkDisplayPowerControlEXT-pDisplayPowerInfo-parameter  
  `pDisplayPowerInfo` **must** be a valid pointer to a valid [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerInfoEXT.html) structure
- [](#VUID-vkDisplayPowerControlEXT-commonparent)VUID-vkDisplayPowerControlEXT-commonparent  
  Both of `device`, and `display` **must** have been created, allocated, or retrieved from the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDisplayPowerControlEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0