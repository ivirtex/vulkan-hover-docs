# VkDeviceEventTypeEXT(3) Manual Page

## Name

VkDeviceEventTypeEXT - Events that can occur on a device object



## [](#_c_specification)C Specification

Possible values of [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceEventInfoEXT.html)::`deviceEvent`, specifying when a fence will be signaled, are:

```c++
// Provided by VK_EXT_display_control
typedef enum VkDeviceEventTypeEXT {
    VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT = 0,
} VkDeviceEventTypeEXT;
```

## [](#_description)Description

- `VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT` specifies that the fence is signaled when a display is plugged into or unplugged from the specified device. Applications **can** use this notification to determine when they need to re-enumerate the available displays on a device.

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceEventInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceEventTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0