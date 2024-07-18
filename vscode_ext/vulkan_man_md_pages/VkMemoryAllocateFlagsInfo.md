# VkMemoryAllocateFlagsInfo(3) Manual Page

## Name

VkMemoryAllocateFlagsInfo - Structure controlling how many instances of
memory will be allocated



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) includes a
`VkMemoryAllocateFlagsInfo` structure, then that structure includes
flags and a device mask controlling how many instances of the memory
will be allocated.

The `VkMemoryAllocateFlagsInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkMemoryAllocateFlagsInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkMemoryAllocateFlags    flags;
    uint32_t                 deviceMask;
} VkMemoryAllocateFlagsInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkMemoryAllocateFlagsInfo VkMemoryAllocateFlagsInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html) controlling
  the allocation.

- `deviceMask` is a mask of physical devices in the logical device,
  indicating that memory **must** be allocated on each device in the
  mask, if `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set in `flags`.

## <a href="#_description" class="anchor"></a>Description

If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is not set, the number of
instances allocated depends on whether
`VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is set in the memory heap. If
`VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is set, then memory is allocated for
every physical device in the logical device (as if `deviceMask` has bits
set for all device indices). If `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` is
not set, then a single instance of memory is allocated (as if
`deviceMask` is set to one).

On some implementations, allocations from a multi-instance heap **may**
consume memory on all physical devices even if the `deviceMask` excludes
some devices. If
[VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)::`subsetAllocation`
is `VK_TRUE`, then memory is only consumed for the devices in the device
mask.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In practice, most allocations on a multi-instance heap will be
allocated across all physical devices. Unicast allocation support is an
optional optimization for a minority of allocations.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkMemoryAllocateFlagsInfo-deviceMask-00675"
  id="VUID-VkMemoryAllocateFlagsInfo-deviceMask-00675"></a>
  VUID-VkMemoryAllocateFlagsInfo-deviceMask-00675  
  If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set, `deviceMask` **must**
  be a valid device mask

- <a href="#VUID-VkMemoryAllocateFlagsInfo-deviceMask-00676"
  id="VUID-VkMemoryAllocateFlagsInfo-deviceMask-00676"></a>
  VUID-VkMemoryAllocateFlagsInfo-deviceMask-00676  
  If `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` is set, `deviceMask` **must**
  not be zero

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryAllocateFlagsInfo-sType-sType"
  id="VUID-VkMemoryAllocateFlagsInfo-sType-sType"></a>
  VUID-VkMemoryAllocateFlagsInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO`

- <a href="#VUID-VkMemoryAllocateFlagsInfo-flags-parameter"
  id="VUID-VkMemoryAllocateFlagsInfo-flags-parameter"></a>
  VUID-VkMemoryAllocateFlagsInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryAllocateFlagsInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
