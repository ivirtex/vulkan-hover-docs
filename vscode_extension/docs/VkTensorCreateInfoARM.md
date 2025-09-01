# VkTensorCreateInfoARM(3) Manual Page

## Name

VkTensorCreateInfoARM - Structure specifying the parameters of a newly created tensor object



## [](#_c_specification)C Specification

The [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorCreateInfoARM {
    VkStructureType                  sType;
    const void*                      pNext;
    VkTensorCreateFlagsARM           flags;
    const VkTensorDescriptionARM*    pDescription;
    VkSharingMode                    sharingMode;
    uint32_t                         queueFamilyIndexCount;
    const uint32_t*                  pQueueFamilyIndices;
} VkTensorCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html) describing additional parameters of the tensor.
- `pDescription` is a pointer to an instance of [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) describing the tensor.
- `sharingMode` is a [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value specifying the sharing mode of the tensor when it will be accessed by multiple queue families.
- `queueFamilyIndexCount` is the number of entries in the `pQueueFamilyIndices` array.
- `pQueueFamilyIndices` is a list of queue families that will access this tensor (ignored if `sharingMode` is not `VK_SHARING_MODE_CONCURRENT`).

## [](#_description)Description

To determine the set of valid `usage` bits for a given tensor format, call [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html) with [VkTensorFormatPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorFormatPropertiesARM.html) in the `pNext` chain.

Tensor Creation Limits

Valid values for some tensor creation parameters are limited by a numerical upper bound or by inclusion in a bitset.

Several limiting values are defined below. The limiting values are referenced by the relevant valid usage statements of `VkTensorCreateInfoARM`.

- Let the uint64\_t tensorElements define the number of data elements in the tensor computed as the product of all `VkTensorCreateInfoARM`::`pDescription->pDimensions`\[i] for i between 0 and `VkTensorCreateInfoARM`::`pDescription->dimensionCount` - 1.

Valid Usage

- [](#VUID-VkTensorCreateInfoARM-pDescription-09720)VUID-VkTensorCreateInfoARM-pDescription-09720  
  If `pDescription->tiling` is `VK_TENSOR_TILING_OPTIMAL_ARM`, `pDescription->pStrides` **must** be `NULL`
- [](#VUID-VkTensorCreateInfoARM-tensorElements-09721)VUID-VkTensorCreateInfoARM-tensorElements-09721  
  `tensorElements` (as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-creation-limits](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-creation-limits)) **must** not be greater than [VkPhysicalDeviceTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorPropertiesARM.html)::`maxTensorElements`
- [](#VUID-VkTensorCreateInfoARM-sharingMode-09722)VUID-VkTensorCreateInfoARM-sharingMode-09722  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, `pQueueFamilyIndices` **must** be a valid pointer to an array of `queueFamilyIndexCount` `uint32_t` values
- [](#VUID-VkTensorCreateInfoARM-sharingMode-09723)VUID-VkTensorCreateInfoARM-sharingMode-09723  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, `queueFamilyIndexCount` **must** be greater than `1`
- [](#VUID-VkTensorCreateInfoARM-sharingMode-09725)VUID-VkTensorCreateInfoARM-sharingMode-09725  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of `pQueueFamilyIndices` **must** be unique and **must** be less than `pQueueFamilyPropertyCount` returned by either [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html) or [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html) for the `physicalDevice` that was used to create `device`
- [](#VUID-VkTensorCreateInfoARM-pNext-09864)VUID-VkTensorCreateInfoARM-pNext-09864  
  If the `pNext` chain includes a [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html) structure, its `handleTypes` member **must** only contain bits that are also in [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html)::`externalMemoryProperties.compatibleHandleTypes`, as returned by [vkGetPhysicalDeviceExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalTensorPropertiesARM.html) with `pExternalTensorInfo->handleType` equal to any one of the handle types specified in [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html)::`handleTypes`
- [](#VUID-VkTensorCreateInfoARM-flags-09726)VUID-VkTensorCreateInfoARM-flags-09726  
  If `flags` includes `VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`, the [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-VkTensorCreateInfoARM-pNext-09727)VUID-VkTensorCreateInfoARM-pNext-09727  
  If the `pNext` chain includes a [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) structure, `flags` **must** contain `VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`
- [](#VUID-VkTensorCreateInfoARM-pDescription-09728)VUID-VkTensorCreateInfoARM-pDescription-09728  
  If `pDescription->usage` does not have any of the following bits set (i.e. if it is not possible to create a tensor view for this tensor), then the [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-view-format-features) **must** contain the format feature flags required by the `usage` flags `pDescription->format` as indicated in the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags) section
  
  - `VK_TENSOR_USAGE_SHADER_BIT_ARM`
  - `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM`
- [](#VUID-VkTensorCreateInfoARM-protectedMemory-09729)VUID-VkTensorCreateInfoARM-protectedMemory-09729  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled, `flags` **must** not contain `VK_TENSOR_CREATE_PROTECTED_BIT_ARM`

Valid Usage (Implicit)

- [](#VUID-VkTensorCreateInfoARM-sType-sType)VUID-VkTensorCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_CREATE_INFO_ARM`
- [](#VUID-VkTensorCreateInfoARM-pNext-pNext)VUID-VkTensorCreateInfoARM-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html) or [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
- [](#VUID-VkTensorCreateInfoARM-sType-unique)VUID-VkTensorCreateInfoARM-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkTensorCreateInfoARM-flags-parameter)VUID-VkTensorCreateInfoARM-flags-parameter  
  `flags` **must** be a valid combination of [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html) values
- [](#VUID-VkTensorCreateInfoARM-pDescription-parameter)VUID-VkTensorCreateInfoARM-pDescription-parameter  
  `pDescription` **must** be a valid pointer to a valid [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure
- [](#VUID-VkTensorCreateInfoARM-sharingMode-parameter)VUID-VkTensorCreateInfoARM-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceTensorMemoryRequirementsARM.html), [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagsARM.html), [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html), [vkCreateTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0