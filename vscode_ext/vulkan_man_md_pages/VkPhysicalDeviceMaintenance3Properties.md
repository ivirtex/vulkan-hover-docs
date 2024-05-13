# VkPhysicalDeviceMaintenance3Properties(3) Manual Page

## Name

VkPhysicalDeviceMaintenance3Properties - Structure describing descriptor
set properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance3Properties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceMaintenance3Properties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxPerSetDescriptors;
    VkDeviceSize       maxMemoryAllocationSize;
} VkPhysicalDeviceMaintenance3Properties;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance3
typedef VkPhysicalDeviceMaintenance3Properties VkPhysicalDeviceMaintenance3PropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-maxPerSetDescriptors"></span>
  `maxPerSetDescriptors` is a maximum number of descriptors (summed over
  all descriptor types) in a single descriptor set that is guaranteed to
  satisfy any implementation-dependent constraints on the size of a
  descriptor set itself. Applications **can** query whether a descriptor
  set that goes beyond this limit is supported using
  [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html).

- <span id="extension-limits-maxMemoryAllocationSize"></span>
  `maxMemoryAllocationSize` is the maximum size of a memory allocation
  that **can** be created, even if there is more space available in the
  heap.

If the `VkPhysicalDeviceMaintenance3Properties` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance3Properties-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance3Properties-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance3Properties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance3Properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
