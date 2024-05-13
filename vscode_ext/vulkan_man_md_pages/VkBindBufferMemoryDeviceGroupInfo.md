# VkBindBufferMemoryDeviceGroupInfo(3) Manual Page

## Name

VkBindBufferMemoryDeviceGroupInfo - Structure specifying device within a
group to bind to



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindBufferMemoryDeviceGroupInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBindBufferMemoryDeviceGroupInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceIndexCount;
    const uint32_t*    pDeviceIndices;
} VkBindBufferMemoryDeviceGroupInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_bind_memory2 with VK_KHR_device_group
typedef VkBindBufferMemoryDeviceGroupInfo VkBindBufferMemoryDeviceGroupInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceIndexCount` is the number of elements in `pDeviceIndices`.

- `pDeviceIndices` is a pointer to an array of device indices.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) includes a
`VkBindBufferMemoryDeviceGroupInfo` structure, then that structure
determines how memory is bound to buffers across multiple devices in a
device group.

If `deviceIndexCount` is greater than zero, then on device index i the
buffer is attached to the instance of `memory` on the physical device
with device index `pDeviceIndices`\[i\].

If `deviceIndexCount` is zero and `memory` comes from a memory heap with
the `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` bit set, then it is as if
`pDeviceIndices` contains consecutive indices from zero to the number of
physical devices in the logical device, minus one. In other words, by
default each physical device attaches to its own instance of `memory`.

If `deviceIndexCount` is zero and `memory` comes from a memory heap
without the `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` bit set, then it is as
if `pDeviceIndices` contains an array of zeros. In other words, by
default each physical device attaches to instance zero.

Valid Usage

- <a href="#VUID-VkBindBufferMemoryDeviceGroupInfo-deviceIndexCount-01606"
  id="VUID-VkBindBufferMemoryDeviceGroupInfo-deviceIndexCount-01606"></a>
  VUID-VkBindBufferMemoryDeviceGroupInfo-deviceIndexCount-01606  
  `deviceIndexCount` **must** either be zero or equal to the number of
  physical devices in the logical device

- <a href="#VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-01607"
  id="VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-01607"></a>
  VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-01607  
  All elements of `pDeviceIndices` **must** be valid device indices

Valid Usage (Implicit)

- <a href="#VUID-VkBindBufferMemoryDeviceGroupInfo-sType-sType"
  id="VUID-VkBindBufferMemoryDeviceGroupInfo-sType-sType"></a>
  VUID-VkBindBufferMemoryDeviceGroupInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO`

- <a
  href="#VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-parameter"
  id="VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-parameter"></a>
  VUID-VkBindBufferMemoryDeviceGroupInfo-pDeviceIndices-parameter  
  If `deviceIndexCount` is not `0`, `pDeviceIndices` **must** be a valid
  pointer to an array of `deviceIndexCount` `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindBufferMemoryDeviceGroupInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
