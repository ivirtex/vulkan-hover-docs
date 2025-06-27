# VkPhysicalDeviceExternalSemaphoreInfo(3) Manual Page

## Name

VkPhysicalDeviceExternalSemaphoreInfo - Structure specifying semaphore
creation parameters.



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExternalSemaphoreInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceExternalSemaphoreInfo {
    VkStructureType                          sType;
    const void*                              pNext;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalSemaphoreInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore_capabilities
typedef VkPhysicalDeviceExternalSemaphoreInfo VkPhysicalDeviceExternalSemaphoreInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value specifying the external semaphore handle type for which
  capabilities will be returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-sType"
  id="VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-sType"></a>
  VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO`

- <a href="#VUID-VkPhysicalDeviceExternalSemaphoreInfo-pNext-pNext"
  id="VUID-VkPhysicalDeviceExternalSemaphoreInfo-pNext-pNext"></a>
  VUID-VkPhysicalDeviceExternalSemaphoreInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)

- <a href="#VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-unique"
  id="VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-unique"></a>
  VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkPhysicalDeviceExternalSemaphoreInfo-handleType-parameter"
  id="VUID-VkPhysicalDeviceExternalSemaphoreInfo-handleType-parameter"></a>
  VUID-VkPhysicalDeviceExternalSemaphoreInfo-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html),
[vkGetPhysicalDeviceExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExternalSemaphoreInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
