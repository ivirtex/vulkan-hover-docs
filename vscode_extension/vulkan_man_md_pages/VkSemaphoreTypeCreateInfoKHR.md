# VkSemaphoreTypeCreateInfo(3) Manual Page

## Name

VkSemaphoreTypeCreateInfo - Structure specifying the type of a newly
created semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreTypeCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreTypeCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphoreType    semaphoreType;
    uint64_t           initialValue;
} VkSemaphoreTypeCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreTypeCreateInfo VkSemaphoreTypeCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphoreType` is a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) value
  specifying the type of the semaphore.

- `initialValue` is the initial payload value if `semaphoreType` is
  `VK_SEMAPHORE_TYPE_TIMELINE`.

## <a href="#_description" class="anchor"></a>Description

To create a semaphore of a specific type, add a
`VkSemaphoreTypeCreateInfo` structure to the
[VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html)::`pNext` chain.

If no `VkSemaphoreTypeCreateInfo` structure is included in the `pNext`
chain of [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html), then the
created semaphore will have a default
[VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`.

Valid Usage

- <a href="#VUID-VkSemaphoreTypeCreateInfo-timelineSemaphore-03252"
  id="VUID-VkSemaphoreTypeCreateInfo-timelineSemaphore-03252"></a>
  VUID-VkSemaphoreTypeCreateInfo-timelineSemaphore-03252  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-timelineSemaphore"
  target="_blank" rel="noopener"><code>timelineSemaphore</code></a>
  feature is not enabled, `semaphoreType` **must** not equal
  `VK_SEMAPHORE_TYPE_TIMELINE`

- <a href="#VUID-VkSemaphoreTypeCreateInfo-semaphoreType-03279"
  id="VUID-VkSemaphoreTypeCreateInfo-semaphoreType-03279"></a>
  VUID-VkSemaphoreTypeCreateInfo-semaphoreType-03279  
  If `semaphoreType` is `VK_SEMAPHORE_TYPE_BINARY`, `initialValue`
  **must** be zero

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreTypeCreateInfo-sType-sType"
  id="VUID-VkSemaphoreTypeCreateInfo-sType-sType"></a>
  VUID-VkSemaphoreTypeCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO`

- <a href="#VUID-VkSemaphoreTypeCreateInfo-semaphoreType-parameter"
  id="VUID-VkSemaphoreTypeCreateInfo-semaphoreType-parameter"></a>
  VUID-VkSemaphoreTypeCreateInfo-semaphoreType-parameter  
  `semaphoreType` **must** be a valid
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreTypeCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
