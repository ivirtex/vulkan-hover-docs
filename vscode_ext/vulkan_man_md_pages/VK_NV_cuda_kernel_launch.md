# VK_NV_cuda_kernel_launch(3) Manual Page

## Name

VK_NV_cuda_kernel_launch - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

308

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

- **This is a *provisional* extension and **must** be used with caution.
  See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-provisional-header"
  target="_blank" rel="noopener">description</a> of provisional header
  files for enablement and stability details.**

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_report

## <a href="#_contact" class="anchor"></a>Contact

- Tristan Lorach <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cuda_kernel_launch%5D%20@tlorach%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cuda_kernel_launch%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tlorach</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-09-30

**Contributors**  
- Eric Werness, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Interoperability between APIs can sometimes create additional overhead
depending on the platform used. This extension targets deployment of
existing CUDA kernels via Vulkan, with a way to directly upload PTX
kernels and dispatch the kernels from Vulkan’s command buffer without
the need to use interoperability between the Vulkan and CUDA contexts.
However, we do encourage actual development using the native CUDA
runtime for the purpose of debugging and profiling.

The application will first have to create a CUDA module using
[vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html) then create the CUDA
function entry point with
[vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaFunctionNV.html).

Then in order to dispatch this function, the application will create a
command buffer where it will launch the kernel with
[vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCudaLaunchKernelNV.html).

When done, the application will then destroy the function handle, as
well as the CUDA module handle with
[vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaFunctionNV.html) and
[vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaModuleNV.html).

To reduce the impact of compilation time, this extension offers the
capability to return a binary cache from the PTX that was provided. For
this, a first query for the required cache size is made with
[vkGetCudaModuleCacheNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCudaModuleCacheNV.html) with a `NULL`
pointer to a buffer and with a valid pointer receiving the size; then
another call of the same function with a valid pointer to a buffer to
retrieve the data. The resulting cache could then be user later for
further runs of this application by sending this cache instead of the
PTX code (using the same
[vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html)), thus significantly
speeding up the initialization of the CUDA module.

As with [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html), the binary cache
depends on the hardware architecture. The application must assume the
cache might fail, and need to handle falling back to the original PTX
code as necessary. Most often, the cache will succeed if the same GPU
driver and architecture is used between the cache generation from PTX
and the use of this cache. In the event of a new driver version, or if
using a different GPU architecture, the cache is likely to become
invalid.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html)

- [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCudaLaunchKernelNV.html)

- [vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaFunctionNV.html)

- [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html)

- [vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaFunctionNV.html)

- [vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaModuleNV.html)

- [vkGetCudaModuleCacheNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetCudaModuleCacheNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCudaFunctionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionCreateInfoNV.html)

- [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaLaunchInfoNV.html)

- [VkCudaModuleCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleCreateInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCudaKernelLaunchFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCudaKernelLaunchFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceCudaKernelLaunchPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCudaKernelLaunchPropertiesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_CUDA_KERNEL_LAUNCH_EXTENSION_NAME`

- `VK_NV_CUDA_KERNEL_LAUNCH_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_CUDA_FUNCTION_NV`

  - `VK_OBJECT_TYPE_CUDA_MODULE_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CUDA_FUNCTION_CREATE_INFO_NV`

  - `VK_STRUCTURE_TYPE_CUDA_LAUNCH_INFO_NV`

  - `VK_STRUCTURE_TYPE_CUDA_MODULE_CREATE_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_PROPERTIES_NV`

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_FUNCTION_NV_EXT`

  - `VK_DEBUG_REPORT_OBJECT_TYPE_CUDA_MODULE_NV_EXT`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-03-01 (Tristan Lorach)

- Revision 2, 2020-09-30 (Tristan Lorach)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_cuda_kernel_launch"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
