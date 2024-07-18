# VkDisplayPowerStateEXT(3) Manual Page

## Name

VkDisplayPowerStateEXT - Possible power states for a display



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerInfoEXT.html)::`powerState`,
specifying the new power state of a display, are:

``` c
// Provided by VK_EXT_display_control
typedef enum VkDisplayPowerStateEXT {
    VK_DISPLAY_POWER_STATE_OFF_EXT = 0,
    VK_DISPLAY_POWER_STATE_SUSPEND_EXT = 1,
    VK_DISPLAY_POWER_STATE_ON_EXT = 2,
} VkDisplayPowerStateEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DISPLAY_POWER_STATE_OFF_EXT` specifies that the display is powered
  down.

- `VK_DISPLAY_POWER_STATE_SUSPEND_EXT` specifies that the display is put
  into a low power mode, from which it **may** be able to transition
  back to `VK_DISPLAY_POWER_STATE_ON_EXT` more quickly than if it were
  in `VK_DISPLAY_POWER_STATE_OFF_EXT`. This state **may** be the same as
  `VK_DISPLAY_POWER_STATE_OFF_EXT`.

- `VK_DISPLAY_POWER_STATE_ON_EXT` specifies that the display is powered
  on.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPowerStateEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
