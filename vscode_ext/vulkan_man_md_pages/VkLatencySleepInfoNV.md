# VkLatencySleepInfoNV(3) Manual Page

## Name

VkLatencySleepInfoNV - Structure specifying the parameters of
vkLatencySleepNV



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkLatencySleepInfoNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkLatencySleepInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphore        signalSemaphore;
    uint64_t           value;
} VkLatencySleepInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `signalSemaphore` is a semaphore that is signaled to indicate that the
  application **should** resume input sampling work.

- `value` is the value that `signalSemaphore` is set to for resuming
  sampling work.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkLatencySleepInfoNV-signalSemaphore-09361"
  id="VUID-VkLatencySleepInfoNV-signalSemaphore-09361"></a>
  VUID-VkLatencySleepInfoNV-signalSemaphore-09361  
  `signalSemaphore` **must** be a timeline semaphore

Valid Usage (Implicit)

- <a href="#VUID-VkLatencySleepInfoNV-sType-sType"
  id="VUID-VkLatencySleepInfoNV-sType-sType"></a>
  VUID-VkLatencySleepInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SLEEP_INFO_NV`

- <a href="#VUID-VkLatencySleepInfoNV-signalSemaphore-parameter"
  id="VUID-VkLatencySleepInfoNV-signalSemaphore-parameter"></a>
  VUID-VkLatencySleepInfoNV-signalSemaphore-parameter  
  `signalSemaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkLatencySleepNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLatencySleepInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
