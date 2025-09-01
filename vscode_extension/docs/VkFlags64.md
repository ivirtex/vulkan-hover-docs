# VkFlags64(3) Manual Page

## Name

VkFlags64 - Vulkan 64-bit bitmasks



## [](#_c_specification)C Specification

A collection of 64-bit flags is represented by a bitmask using the type [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html):

```c++
// Provided by VK_VERSION_1_3, VK_KHR_synchronization2
typedef uint64_t VkFlags64;
```

## [](#_description)Description

When the 31 bits available in [VkFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags.html) are insufficient, the [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html) type can be passed to commands and structures to represent up to 64 options. [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html) is not used directly in the API. Instead, a `Vk*Flags2` type which is an alias of [VkFlags64](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags64.html), and whose name matches the corresponding `Vk*FlagBits2` that are valid for that type, is used.

Any `Vk*Flags2` member or parameter used in the API as an input **must** be a valid combination of bit flags. A valid combination is either zero or the bitwise OR of valid bit flags.

An individual bit flag is valid for a `Vk*Flags2` type if it would be a [valid enumerant](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-validusage-enums) when used with the equivalent `Vk*FlagBits2` type, where the bits type is obtained by taking the flag type and replacing the trailing `Flags2` with `FlagBits2`. For example, a flag value of type [VkAccessFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2KHR.html) **must** contain only bit flags defined by [VkAccessFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2KHR.html).

Any `Vk*Flags2` member or parameter returned from a query command or otherwise output from Vulkan to the application **may** contain bit flags undefined in its corresponding `Vk*FlagBits2` type. An application **cannot** rely on the state of these unspecified bits.

Note

Both the `Vk*FlagBits2` type, and the individual bits defined for that type, are defined as `uint64_t` integers in the C API. This is in contrast to the 32-bit types, where the `Vk*FlagBits` type is defined as a C `enum` and the individual bits as enumerants belonging to that `enum`. As a result, there is less compile time type checking possible for the 64-bit types. This is unavoidable since there is no sufficiently portable way to define a 64-bit `enum` type in C99.

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html), [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html), [VkBufferUsageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2.html), [VkDataGraphPipelineDispatchFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchFlagsARM.html), [VkDataGraphPipelineSessionCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagsARM.html), [VkFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags.html), [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html), [VkMemoryDecompressionMethodFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDecompressionMethodFlagsNV.html), [VkPhysicalDeviceSchedulingControlsFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSchedulingControlsFlagsARM.html), [VkPipelineCreateFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkTensorCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagsARM.html), [VkTensorUsageFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagsARM.html), [VkTensorViewCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFlags64).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0