# VkDeviceEventTypeEXT(3) Manual Page

## Name

VkDeviceEventTypeEXT - Events that can occur on a device object



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventInfoEXT.html)::`device`, specifying
when a fence will be signaled, are:

``` c
// Provided by VK_EXT_display_control
typedef enum VkDeviceEventTypeEXT {
    VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT = 0,
} VkDeviceEventTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT` specifies that the fence is
  signaled when a display is plugged into or unplugged from the
  specified device. Applications **can** use this notification to
  determine when they need to re-enumerate the available displays on a
  device.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceEventTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
