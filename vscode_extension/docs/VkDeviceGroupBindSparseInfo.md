# VkDeviceGroupBindSparseInfo(3) Manual Page

## Name

VkDeviceGroupBindSparseInfo - Structure indicating which instances are bound



## [](#_c_specification)C Specification

If the `pNext` chain of [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html) includes a `VkDeviceGroupBindSparseInfo` structure, then that structure includes device indices specifying which instance of the resources and memory are bound.

The `VkDeviceGroupBindSparseInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupBindSparseInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           resourceDeviceIndex;
    uint32_t           memoryDeviceIndex;
} VkDeviceGroupBindSparseInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group
typedef VkDeviceGroupBindSparseInfo VkDeviceGroupBindSparseInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `resourceDeviceIndex` is a device index indicating which instance of the resource is bound.
- `memoryDeviceIndex` is a device index indicating which instance of the memory the resource instance is bound to.

## [](#_description)Description

These device indices apply to all buffer and image memory binds included in the batch pointing to this structure. The semaphore waits and signals for the batch are executed only by the physical device specified by the `resourceDeviceIndex`.

If this structure is not present, `resourceDeviceIndex` and `memoryDeviceIndex` are assumed to be zero.

Valid Usage

- [](#VUID-VkDeviceGroupBindSparseInfo-resourceDeviceIndex-01118)VUID-VkDeviceGroupBindSparseInfo-resourceDeviceIndex-01118  
  `resourceDeviceIndex` and `memoryDeviceIndex` **must** both be valid device indices
- [](#VUID-VkDeviceGroupBindSparseInfo-memoryDeviceIndex-01119)VUID-VkDeviceGroupBindSparseInfo-memoryDeviceIndex-01119  
  Each memory allocation bound in this batch **must** have allocated an instance for `memoryDeviceIndex`

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupBindSparseInfo-sType-sType)VUID-VkDeviceGroupBindSparseInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupBindSparseInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0