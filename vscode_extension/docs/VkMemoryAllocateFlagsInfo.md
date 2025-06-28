# VkMemoryAllocateFlagsInfo(3) Manual Page

## Name

VkMemoryAllocateFlagsInfo - Structure controlling how many instances of memory will be allocated



## [](#_c_specification)C Specification

If the `pNext` chain of [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) includes a `VkMemoryAllocateFlagsInfo` structure, then that structure includes flags and a device mask controlling how many instances of the memory will be allocated.

The `VkMemoryAllocateFlagsInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkMemoryAllocateFlagsInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkMemoryAllocateFlags    flags;
    uint32_t                 deviceMask;
} VkMemoryAllocateFlagsInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group
typedef VkMemoryAllocateFlagsInfo VkMemoryAllocateFlagsInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html) controlling the allocation.
- `deviceMask` is a mask of physical devices in the logical device, indicating that memory **must** be allocated on each device in the mask, if `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set in `flags`.

## [](#_description)Description

If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is not set, the number of instances allocated depends on whether `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is set in the memory heap. If `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is set, then memory is allocated for every physical device in the logical device (as if `deviceMask` has bits set for all device indices). If `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is not set, then a single instance of memory is allocated (as if `deviceMask` is set to one).

On some implementations, allocations from a multi-instance heap **may** consume memory on all physical devices even if the `deviceMask` excludes some devices. If [VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGroupProperties.html)::`subsetAllocation` is `VK_TRUE`, then memory is only consumed for the devices in the device mask.

Note

In practice, most allocations on a multi-instance heap will be allocated across all physical devices. Unicast allocation support is an optional optimization for a minority of allocations.

Valid Usage

- [](#VUID-VkMemoryAllocateFlagsInfo-deviceMask-00675)VUID-VkMemoryAllocateFlagsInfo-deviceMask-00675  
  If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set, `deviceMask` **must** be a valid device mask
- [](#VUID-VkMemoryAllocateFlagsInfo-deviceMask-00676)VUID-VkMemoryAllocateFlagsInfo-deviceMask-00676  
  If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set, `deviceMask` **must** not be zero
- [](#VUID-VkMemoryAllocateFlagsInfo-flags-10760)VUID-VkMemoryAllocateFlagsInfo-flags-10760  
  If the allocation is performing a memory import operation, then `flags` **must** not contain `VK_MEMORY_ALLOCATE_ZERO_INITIALIZE_BIT_EXT`
- [](#VUID-VkMemoryAllocateFlagsInfo-flags-10761)VUID-VkMemoryAllocateFlagsInfo-flags-10761  
  If the allocation uses protected memory, then `flags` **must** not contain `VK_MEMORY_ALLOCATE_ZERO_INITIALIZE_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkMemoryAllocateFlagsInfo-sType-sType)VUID-VkMemoryAllocateFlagsInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO`
- [](#VUID-VkMemoryAllocateFlagsInfo-flags-parameter)VUID-VkMemoryAllocateFlagsInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryAllocateFlagsInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0