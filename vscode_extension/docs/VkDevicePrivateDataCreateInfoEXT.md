# VkDevicePrivateDataCreateInfo(3) Manual Page

## Name

VkDevicePrivateDataCreateInfo - Reserve private data slots



## [](#_c_specification)C Specification

To reserve private data storage slots, add a [VkDevicePrivateDataCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePrivateDataCreateInfo.html) structure to the `pNext` chain of the [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) structure. Reserving slots in this manner is not strictly necessary, but doing so **may** improve performance.

```c++
// Provided by VK_VERSION_1_3
typedef struct VkDevicePrivateDataCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           privateDataSlotRequestCount;
} VkDevicePrivateDataCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_private_data
typedef VkDevicePrivateDataCreateInfo VkDevicePrivateDataCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `privateDataSlotRequestCount` is the amount of slots to reserve.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDevicePrivateDataCreateInfo-sType-sType)VUID-VkDevicePrivateDataCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO`

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDevicePrivateDataCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0