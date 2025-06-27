# VkDeviceGroupCommandBufferBeginInfo(3) Manual Page

## Name

VkDeviceGroupCommandBufferBeginInfo - Set the initial device mask for a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html) includes a
`VkDeviceGroupCommandBufferBeginInfo` structure, then that structure
includes an initial device mask for the command buffer.

The `VkDeviceGroupCommandBufferBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupCommandBufferBeginInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceMask;
} VkDeviceGroupCommandBufferBeginInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkDeviceGroupCommandBufferBeginInfo VkDeviceGroupCommandBufferBeginInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceMask` is the initial value of the command buffer’s device mask.

## <a href="#_description" class="anchor"></a>Description

The initial device mask also acts as an upper bound on the set of
devices that **can** ever be in the device mask in the command buffer.

If this structure is not present, the initial value of a command
buffer’s device mask is set to include all physical devices in the
logical device when the command buffer begins recording.

Valid Usage

- <a href="#VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00106"
  id="VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00106"></a>
  VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00106  
  `deviceMask` **must** be a valid device mask value

- <a href="#VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00107"
  id="VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00107"></a>
  VUID-VkDeviceGroupCommandBufferBeginInfo-deviceMask-00107  
  `deviceMask` **must** not be zero

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupCommandBufferBeginInfo-sType-sType"
  id="VUID-VkDeviceGroupCommandBufferBeginInfo-sType-sType"></a>
  VUID-VkDeviceGroupCommandBufferBeginInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupCommandBufferBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
