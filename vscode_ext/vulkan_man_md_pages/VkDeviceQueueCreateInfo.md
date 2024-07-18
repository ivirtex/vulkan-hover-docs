# VkDeviceQueueCreateInfo(3) Manual Page

## Name

VkDeviceQueueCreateInfo - Structure specifying parameters of a newly
created device queue



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceQueueCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDeviceQueueCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceQueueCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
    uint32_t                    queueCount;
    const float*                pQueuePriorities;
} VkDeviceQueueCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask indicating behavior of the queues.

- `queueFamilyIndex` is an unsigned integer indicating the index of the
  queue family in which to create the queues on this device. This index
  corresponds to the index of an element of the `pQueueFamilyProperties`
  array that was returned by `vkGetPhysicalDeviceQueueFamilyProperties`.

- `queueCount` is an unsigned integer specifying the number of queues to
  create in the queue family indicated by `queueFamilyIndex`, and with
  the behavior specified by `flags`.

- `pQueuePriorities` is a pointer to an array of `queueCount` normalized
  floating-point values, specifying priorities of work that will be
  submitted to each created queue. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-priority"
  target="_blank" rel="noopener">Queue Priority</a> for more
  information.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDeviceQueueCreateInfo-queueFamilyIndex-00381"
  id="VUID-VkDeviceQueueCreateInfo-queueFamilyIndex-00381"></a>
  VUID-VkDeviceQueueCreateInfo-queueFamilyIndex-00381  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties`

- <a href="#VUID-VkDeviceQueueCreateInfo-queueCount-00382"
  id="VUID-VkDeviceQueueCreateInfo-queueCount-00382"></a>
  VUID-VkDeviceQueueCreateInfo-queueCount-00382  
  `queueCount` **must** be less than or equal to the `queueCount` member
  of the `VkQueueFamilyProperties` structure, as returned by
  `vkGetPhysicalDeviceQueueFamilyProperties` in the
  `pQueueFamilyProperties`\[queueFamilyIndex\]

- <a href="#VUID-VkDeviceQueueCreateInfo-pQueuePriorities-00383"
  id="VUID-VkDeviceQueueCreateInfo-pQueuePriorities-00383"></a>
  VUID-VkDeviceQueueCreateInfo-pQueuePriorities-00383  
  Each element of `pQueuePriorities` **must** be between `0.0` and `1.0`
  inclusive

- <a href="#VUID-VkDeviceQueueCreateInfo-flags-02861"
  id="VUID-VkDeviceQueueCreateInfo-flags-02861"></a>
  VUID-VkDeviceQueueCreateInfo-flags-02861  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled, the `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT` bit
  of `flags` **must** not be set

- <a href="#VUID-VkDeviceQueueCreateInfo-flags-06449"
  id="VUID-VkDeviceQueueCreateInfo-flags-06449"></a>
  VUID-VkDeviceQueueCreateInfo-flags-06449  
  If `flags` includes `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`,
  `queueFamilyIndex` **must** be the index of a queue family that
  includes the `VK_QUEUE_PROTECTED_BIT` capability

- <a href="#VUID-VkDeviceQueueCreateInfo-pNext-09398"
  id="VUID-VkDeviceQueueCreateInfo-pNext-09398"></a>
  VUID-VkDeviceQueueCreateInfo-pNext-09398  
  If the `pNext` chain includes a
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)
  structure then
  [VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html)::`schedulingControlsFlags`
  **must** contain
  `VK_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_SHADER_CORE_COUNT_ARM`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceQueueCreateInfo-sType-sType"
  id="VUID-VkDeviceQueueCreateInfo-sType-sType"></a>
  VUID-VkDeviceQueueCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO`

- <a href="#VUID-VkDeviceQueueCreateInfo-pNext-pNext"
  id="VUID-VkDeviceQueueCreateInfo-pNext-pNext"></a>
  VUID-VkDeviceQueueCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html)
  or
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)

- <a href="#VUID-VkDeviceQueueCreateInfo-sType-unique"
  id="VUID-VkDeviceQueueCreateInfo-sType-unique"></a>
  VUID-VkDeviceQueueCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkDeviceQueueCreateInfo-flags-parameter"
  id="VUID-VkDeviceQueueCreateInfo-flags-parameter"></a>
  VUID-VkDeviceQueueCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html) values

- <a href="#VUID-VkDeviceQueueCreateInfo-pQueuePriorities-parameter"
  id="VUID-VkDeviceQueueCreateInfo-pQueuePriorities-parameter"></a>
  VUID-VkDeviceQueueCreateInfo-pQueuePriorities-parameter  
  `pQueuePriorities` **must** be a valid pointer to an array of
  `queueCount` `float` values

- <a href="#VUID-VkDeviceQueueCreateInfo-queueCount-arraylength"
  id="VUID-VkDeviceQueueCreateInfo-queueCount-arraylength"></a>
  VUID-VkDeviceQueueCreateInfo-queueCount-arraylength  
  `queueCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html),
[VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceQueueCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
