# VkDevicePrivateDataCreateInfo(3) Manual Page

## Name

VkDevicePrivateDataCreateInfo - Reserve private data slots



## <a href="#_c_specification" class="anchor"></a>C Specification

To reserve private data storage slots, add a
[VkDevicePrivateDataCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevicePrivateDataCreateInfo.html)
structure to the `pNext` chain of the
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure. Reserving slots
in this manner is not strictly necessary, but doing so **may** improve
performance.

``` c
// Provided by VK_VERSION_1_3
typedef struct VkDevicePrivateDataCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           privateDataSlotRequestCount;
} VkDevicePrivateDataCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_private_data
typedef VkDevicePrivateDataCreateInfo VkDevicePrivateDataCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `privateDataSlotRequestCount` is the amount of slots to reserve.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDevicePrivateDataCreateInfo-sType-sType"
  id="VUID-VkDevicePrivateDataCreateInfo-sType-sType"></a>
  VUID-VkDevicePrivateDataCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_private_data](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_private_data.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDevicePrivateDataCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
