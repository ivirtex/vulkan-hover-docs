# VkDeviceGroupCommandBufferBeginInfo(3) Manual Page

## Name

VkDeviceGroupCommandBufferBeginInfo - Set the initial device mask for a command buffer



## [](#_c_specification)C Specification

If the `pNext` chain of [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html) includes a `VkDeviceGroupCommandBufferBeginInfo` structure, then that structure includes an initial device mask for the command buffer.

The `VkDeviceGroupCommandBufferBeginInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupCommandBufferBeginInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceMask;
} VkDeviceGroupCommandBufferBeginInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group
typedef VkDeviceGroupCommandBufferBeginInfo VkDeviceGroupCommandBufferBeginInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `deviceMask` is the initial value of the command buffer’s device mask.

## [](#_description)Description

The initial device mask also acts as an upper bound on the set of devices that **can** ever be in the device mask in the command buffer.

If this structure is not present, the initial value of a command buffer’s device mask includes all physical devices in the logical device when the command buffer begins recording.

Valid Usage

- [](#VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00106)VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00106  
  `deviceMask` **must** be a valid device mask value
- [](#VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00107)VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00107  
  `deviceMask` **must** not be zero

Valid Usage (Implicit)

- [](#VUID-VkDeviceGroupCommandBufferBeginInfo-sType-sType)VUID-VkDeviceGroupCommandBufferBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceGroupCommandBufferBeginInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0