# VkPhysicalDeviceSchedulingControlsFlagBitsARM(3) Manual Page

## Name

VkPhysicalDeviceSchedulingControlsFlagBitsARM - Bitmask specifying
scheduling controls supported by a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html)::`schedulingControlsFlags`,
specifying supported scheduling controls, are:

``` c
// Provided by VK_ARM_scheduling_controls
// Flag bits for VkPhysicalDeviceSchedulingControlsFlagBitsARM
typedef VkFlags64 VkPhysicalDeviceSchedulingControlsFlagBitsARM;
static const VkPhysicalDeviceSchedulingControlsFlagBitsARM VK_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_SHADER_CORE_COUNT_ARM = 0x00000001ULL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_SHADER_CORE_COUNT_ARM`
  indicates that a
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)
  structure **may** be included in the `pNext` chain of a
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) or
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_scheduling_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_scheduling_controls.html),
[VkPhysicalDeviceSchedulingControlsFlagsARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFlagsARM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSchedulingControlsFlagBitsARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
