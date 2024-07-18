# VkSemaphoreSignalInfo(3) Manual Page

## Name

VkSemaphoreSignalInfo - Structure containing information about a
semaphore signal operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreSignalInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreSignalInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphore        semaphore;
    uint64_t           value;
} VkSemaphoreSignalInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreSignalInfo VkSemaphoreSignalInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphore` is the handle of the semaphore to signal.

- `value` is the value to signal.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSemaphoreSignalInfo-semaphore-03257"
  id="VUID-VkSemaphoreSignalInfo-semaphore-03257"></a>
  VUID-VkSemaphoreSignalInfo-semaphore-03257  
  `semaphore` **must** have been created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`

- <a href="#VUID-VkSemaphoreSignalInfo-value-03258"
  id="VUID-VkSemaphoreSignalInfo-value-03258"></a>
  VUID-VkSemaphoreSignalInfo-value-03258  
  `value` **must** have a value greater than the current value of the
  semaphore

- <a href="#VUID-VkSemaphoreSignalInfo-value-03259"
  id="VUID-VkSemaphoreSignalInfo-value-03259"></a>
  VUID-VkSemaphoreSignalInfo-value-03259  
  `value` **must** be less than the value of any pending semaphore
  signal operations

- <a href="#VUID-VkSemaphoreSignalInfo-value-03260"
  id="VUID-VkSemaphoreSignalInfo-value-03260"></a>
  VUID-VkSemaphoreSignalInfo-value-03260  
  `value` **must** have a value which does not differ from the current
  value of the semaphore or the value of any outstanding semaphore wait
  or signal operation on `semaphore` by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreSignalInfo-sType-sType"
  id="VUID-VkSemaphoreSignalInfo-sType-sType"></a>
  VUID-VkSemaphoreSignalInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO`

- <a href="#VUID-VkSemaphoreSignalInfo-pNext-pNext"
  id="VUID-VkSemaphoreSignalInfo-pNext-pNext"></a>
  VUID-VkSemaphoreSignalInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSemaphoreSignalInfo-semaphore-parameter"
  id="VUID-VkSemaphoreSignalInfo-semaphore-parameter"></a>
  VUID-VkSemaphoreSignalInfo-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html),
[vkSignalSemaphoreKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphoreKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreSignalInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
