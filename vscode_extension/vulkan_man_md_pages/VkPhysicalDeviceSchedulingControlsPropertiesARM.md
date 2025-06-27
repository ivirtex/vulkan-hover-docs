# VkPhysicalDeviceSchedulingControlsPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceSchedulingControlsPropertiesARM - Structure containing
scheduling control properties of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSchedulingControlsPropertiesARM` structure is
defined as:

``` c
// Provided by VK_ARM_scheduling_controls
typedef struct VkPhysicalDeviceSchedulingControlsPropertiesARM {
    VkStructureType                               sType;
    void*                                         pNext;
    VkPhysicalDeviceSchedulingControlsFlagsARM    schedulingControlsFlags;
} VkPhysicalDeviceSchedulingControlsPropertiesARM;
```

## <a href="#_members" class="anchor"></a>Members

- <span id="limits-schedulingControlsFlags"></span>`schedulingControlsFlags`
  specifies the specific scheduling controls that a physical device
  supports.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceSchedulingControlsPropertiesARM` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-sType-sType"
  id="VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-sType-sType"></a>
  VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_PROPERTIES_ARM`

- <a
  href="#VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-parameter"
  id="VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-parameter"></a>
  VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-parameter  
  `schedulingControlsFlags` **must** be a valid combination of
  [VkPhysicalDeviceSchedulingControlsFlagBitsARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFlagBitsARM.html)
  values

- <a
  href="#VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-requiredbitmask"
  id="VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-requiredbitmask"></a>
  VUID-VkPhysicalDeviceSchedulingControlsPropertiesARM-schedulingControlsFlags-requiredbitmask  
  `schedulingControlsFlags` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ARM_scheduling_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ARM_scheduling_controls.html),
[VkPhysicalDeviceSchedulingControlsFlagsARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFlagsARM.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSchedulingControlsPropertiesARM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
