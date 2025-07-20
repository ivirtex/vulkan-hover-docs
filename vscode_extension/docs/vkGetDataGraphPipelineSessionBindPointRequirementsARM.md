# vkGetDataGraphPipelineSessionBindPointRequirementsARM(3) Manual Page

## Name

vkGetDataGraphPipelineSessionBindPointRequirementsARM - Get the bind point requirements of a data graph pipeline session



## [](#_c_specification)C Specification

To determine the bind point requirements for a data graph pipeline session, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkGetDataGraphPipelineSessionBindPointRequirementsARM(
    VkDevice                                    device,
    const VkDataGraphPipelineSessionBindPointRequirementsInfoARM* pInfo,
    uint32_t*                                   pBindPointRequirementCount,
    VkDataGraphPipelineSessionBindPointRequirementARM* pBindPointRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the data graph pipeline session.
- `pInfo` is a pointer to a [VkDataGraphPipelineSessionBindPointRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementsInfoARM.html) structure containing parameters for the bind point requirements query.
- `pBindPointRequirementCount` is a pointer to an integer related to the number of bind point available or queried, as described below.
- `pBindPointRequirements` is either `NULL` or a pointer to an array of [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html) structures.

## [](#_description)Description

If `pBindPointRequirements` is `NULL`, then the number of bind points associated with the data graph pipeline session is returned in `pBindPointRequirementCount`. Otherwise, `pBindPointRequirementCount` **must** point to a variable set by the user to the number of elements in the `pBindPointRequirements` array, and on return the variable is overwritten with the number of structures actually written to `pBindPointRequirements`. If `pBindPointRequirementCount` is less than the number of bind points associated with the data graph pipeline session, at most `pBindPointRequirementCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the required bind points were returned.

Valid Usage

- [](#VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-session-09783)VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-session-09783  
  The `session` member of `pInfo` **must** have been created with `device`

Valid Usage (Implicit)

- [](#VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-device-parameter)VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pInfo-parameter)VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineSessionBindPointRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementsInfoARM.html) structure
- [](#VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pBindPointRequirementCount-parameter)VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pBindPointRequirementCount-parameter  
  `pBindPointRequirementCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pBindPointRequirements-parameter)VUID-vkGetDataGraphPipelineSessionBindPointRequirementsARM-pBindPointRequirements-parameter  
  If the value referenced by `pBindPointRequirementCount` is not `0`, and `pBindPointRequirements` is not `NULL`, `pBindPointRequirements` **must** be a valid pointer to an array of `pBindPointRequirementCount` [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html), [VkDataGraphPipelineSessionBindPointRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementsInfoARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDataGraphPipelineSessionBindPointRequirementsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0