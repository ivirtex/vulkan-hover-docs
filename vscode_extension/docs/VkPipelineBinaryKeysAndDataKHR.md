# VkPipelineBinaryKeysAndDataKHR(3) Manual Page

## Name

VkPipelineBinaryKeysAndDataKHR - Structure specifying arrays of key and data pairs



## [](#_c_specification)C Specification

The `VkPipelineBinaryKeysAndDataKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryKeysAndDataKHR {
    uint32_t                          binaryCount;
    const VkPipelineBinaryKeyKHR*     pPipelineBinaryKeys;
    const VkPipelineBinaryDataKHR*    pPipelineBinaryData;
} VkPipelineBinaryKeysAndDataKHR;
```

## [](#_members)Members

- `binaryCount` is the size of the `pPipelineBinaryKeys` and `pPipelineBinaryData` arrays
- `pPipelineBinaryKeys` is a pointer to an array of [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structures containing the pipeline binary keys
- `pPipelineBinaryData` is a pointer to an array of [VkPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataKHR.html) structures containing the pipeline binary data

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryKeysAndDataKHR-pPipelineBinaryKeys-parameter)VUID-VkPipelineBinaryKeysAndDataKHR-pPipelineBinaryKeys-parameter  
  `pPipelineBinaryKeys` **must** be a valid pointer to an array of `binaryCount` valid [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structures
- [](#VUID-VkPipelineBinaryKeysAndDataKHR-pPipelineBinaryData-parameter)VUID-VkPipelineBinaryKeysAndDataKHR-pPipelineBinaryData-parameter  
  `pPipelineBinaryData` **must** be a valid pointer to an array of `binaryCount` valid [VkPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataKHR.html) structures
- [](#VUID-VkPipelineBinaryKeysAndDataKHR-binaryCount-arraylength)VUID-VkPipelineBinaryKeysAndDataKHR-binaryCount-arraylength  
  `binaryCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html), [VkPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataKHR.html), [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryKeysAndDataKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0