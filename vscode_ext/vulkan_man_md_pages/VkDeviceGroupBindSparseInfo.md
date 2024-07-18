# VkDeviceGroupBindSparseInfo(3) Manual Page

## Name

VkDeviceGroupBindSparseInfo - Structure indicating which instances are
bound



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html)
includes a `VkDeviceGroupBindSparseInfo` structure, then that structure
includes device indices specifying which instance of the resources and
memory are bound.

The `VkDeviceGroupBindSparseInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupBindSparseInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           resourceDeviceIndex;
    uint32_t           memoryDeviceIndex;
} VkDeviceGroupBindSparseInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkDeviceGroupBindSparseInfo VkDeviceGroupBindSparseInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `resourceDeviceIndex` is a device index indicating which instance of
  the resource is bound.

- `memoryDeviceIndex` is a device index indicating which instance of the
  memory the resource instance is bound to.

## <a href="#_description" class="anchor"></a>Description

These device indices apply to all buffer and image memory binds included
in the batch pointing to this structure. The semaphore waits and signals
for the batch are executed only by the physical device specified by the
`resourceDeviceIndex`.

If this structure is not present, `resourceDeviceIndex` and
`memoryDeviceIndex` are assumed to be zero.

Valid Usage

- <a href="#VUID-VkDeviceGroupBindSparseInfo-resourceDeviceIndex-01118"
  id="VUID-VkDeviceGroupBindSparseInfo-resourceDeviceIndex-01118"></a>
  VUID-VkDeviceGroupBindSparseInfo-resourceDeviceIndex-01118  
  `resourceDeviceIndex` and `memoryDeviceIndex` **must** both be valid
  device indices

- <a href="#VUID-VkDeviceGroupBindSparseInfo-memoryDeviceIndex-01119"
  id="VUID-VkDeviceGroupBindSparseInfo-memoryDeviceIndex-01119"></a>
  VUID-VkDeviceGroupBindSparseInfo-memoryDeviceIndex-01119  
  Each memory allocation bound in this batch **must** have allocated an
  instance for `memoryDeviceIndex`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupBindSparseInfo-sType-sType"
  id="VUID-VkDeviceGroupBindSparseInfo-sType-sType"></a>
  VUID-VkDeviceGroupBindSparseInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupBindSparseInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
