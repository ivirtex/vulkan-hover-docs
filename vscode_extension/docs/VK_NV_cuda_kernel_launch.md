# VK\_NV\_cuda\_kernel\_launch(3) Manual Page

## Name

VK\_NV\_cuda\_kernel\_launch - device extension



## [](#_registered_extension_number)Registered Extension Number

308

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

- **This is a *provisional* extension and must be used with caution. See the [description](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-provisional-header) of provisional header files for enablement and stability details.**

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_report

## [](#_contact)Contact

- Tristan Lorach [\[GitHub\]tlorach](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cuda_kernel_launch%5D%20%40tlorach%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cuda_kernel_launch%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-09-30

**Contributors**

- Eric Werness, NVIDIA

## [](#_description)Description

Interoperability between APIs can sometimes create additional overhead depending on the platform used. This extension targets deployment of existing CUDA kernels via Vulkan, with a way to directly upload PTX kernels and dispatch the kernels from Vulkan’s command buffer without the need to use interoperability between the Vulkan and CUDA contexts. However, we do encourage actual development using the native CUDA runtime for the purpose of debugging and profiling.

The application will first have to create a CUDA module using [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html) then create the CUDA function entry point with [vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaFunctionNV.html).

Then in order to dispatch this function, the application will create a command buffer where it will launch the kernel with [vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCudaLaunchKernelNV.html).

When done, the application will then destroy the function handle, as well as the CUDA module handle with [vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaFunctionNV.html) and [vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaModuleNV.html).

To reduce the impact of compilation time, this extension offers the capability to return a binary cache from the PTX that was provided. For this, a first query for the required cache size is made with [vkGetCudaModuleCacheNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCudaModuleCacheNV.html) with a `NULL` pointer to a buffer and with a valid pointer receiving the size; then another call of the same function with a valid pointer to a buffer to retrieve the data. The resulting cache could then be used later for further runs of this application by sending this cache instead of the PTX code (using the same [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html)), thus significantly speeding up the initialization of the CUDA module.

As with [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html), the binary cache depends on the hardware architecture. The application must assume the cache might fail, and need to handle falling back to the original PTX code as necessary. Most often, the cache will succeed if the same GPU driver and architecture is used between the cache generation from PTX and the use of this cache. In the event of a new driver version, or if using a different GPU architecture, the cache is likely to become invalid.

## [](#_new_object_types)New Object Types

- [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html)
- [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html)

## [](#_new_commands)New Commands

- [vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCudaLaunchKernelNV.html)
- [vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaFunctionNV.html)
- [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html)
- [vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaFunctionNV.html)
- [vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCudaModuleNV.html)
- [vkGetCudaModuleCacheNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetCudaModuleCacheNV.html)

## [](#_new_structures)New Structures

- [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionCreateInfoNV.html)
- [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaLaunchInfoNV.html)
- [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleCreateInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCudaKernelLaunchFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCudaKernelLaunchFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCudaKernelLaunchPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCudaKernelLaunchPropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_CUDA_KERNEL_LAUNCH_EXTENSION_NAME`
- `VK_NV_CUDA_KERNEL_LAUNCH_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_CUDA_FUNCTION_NV`
  - `VK_OBJECT_TYPE_CUDA_MODULE_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CUDA_FUNCTION_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_CUDA_LAUNCH_INFO_NV`
  - `VK_STRUCTURE_TYPE_CUDA_MODULE_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_PROPERTIES_NV`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_FUNCTION_NV_EXT`
  - `VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_MODULE_NV_EXT`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2020-03-01 (Tristan Lorach)
- Revision 2, 2020-09-30 (Tristan Lorach)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_cuda_kernel_launch).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0