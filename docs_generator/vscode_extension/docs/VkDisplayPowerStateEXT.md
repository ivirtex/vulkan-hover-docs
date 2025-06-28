# VkDisplayPowerStateEXT(3) Manual Page

## Name

VkDisplayPowerStateEXT - Possible power states for a display



## [](#_c_specification)C Specification

Possible values of [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerInfoEXT.html)::`powerState`, specifying the new power state of a display, are:

```c++
// Provided by VK_EXT_display_control
typedef enum VkDisplayPowerStateEXT {
    VK_DISPLAY_POWER_STATE_OFF_EXT = 0,
    VK_DISPLAY_POWER_STATE_SUSPEND_EXT = 1,
    VK_DISPLAY_POWER_STATE_ON_EXT = 2,
} VkDisplayPowerStateEXT;
```

## [](#_description)Description

- `VK_DISPLAY_POWER_STATE_OFF_EXT` specifies that the display is powered down.
- `VK_DISPLAY_POWER_STATE_SUSPEND_EXT` specifies that the display is put into a low power mode, from which it **may** be able to transition back to `VK_DISPLAY_POWER_STATE_ON_EXT` more quickly than if it were in `VK_DISPLAY_POWER_STATE_OFF_EXT`. This state **may** be the same as `VK_DISPLAY_POWER_STATE_OFF_EXT`.
- `VK_DISPLAY_POWER_STATE_ON_EXT` specifies that the display is powered on.

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayPowerStateEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0