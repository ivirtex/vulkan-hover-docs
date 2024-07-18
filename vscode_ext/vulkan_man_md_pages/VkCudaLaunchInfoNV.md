# VkCudaLaunchInfoNV(3) Manual Page

## Name

VkCudaLaunchInfoNV - Structure specifying the parameters to launch a
CUDA kernel



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCudaLaunchInfoNV` structure is very close to the parameters of
the CUDA-Driver function
[cuLaunchKernel](https://docs.nvidia.com/cuda/cuda-driver-api/group__CUDA__EXEC.html#group__CUDA__EXEC_1gb8f3dc3031b40da29d5f9a7139e52e15)
documented in section [6.19 Execution
Control](https://docs.nvidia.com/cuda/cuda-driver-api/group__CUDA__EXEC.html#group__CUDA__EXEC)
of CUDA Driver API.

The structure is defined as:

``` c
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkCudaLaunchInfoNV {
    VkStructureType        sType;
    const void*            pNext;
    VkCudaFunctionNV       function;
    uint32_t               gridDimX;
    uint32_t               gridDimY;
    uint32_t               gridDimZ;
    uint32_t               blockDimX;
    uint32_t               blockDimY;
    uint32_t               blockDimZ;
    uint32_t               sharedMemBytes;
    size_t                 paramCount;
    const void* const *    pParams;
    size_t                 extraCount;
    const void* const *    pExtras;
} VkCudaLaunchInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `function` is the CUDA-Driver handle to the function being launched.

- `gridDimX` is the number of local workgroups to dispatch in the X
  dimension. It must be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\]

- `gridDimY` is the number of local workgroups to dispatch in the Y
  dimension. It must be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\]

- `gridDimZ` is the number of local workgroups to dispatch in the Z
  dimension. It must be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\]

- `blockDimX` is block size in the X dimension.

- `blockDimY` is block size in the Y dimension.

- `blockDimZ` is block size in the Z dimension.

- `sharedMemBytes` is the dynamic shared-memory size per thread block in
  bytes.

- `paramCount` is the length of the `pParams` table.

- `pParams` is a pointer to an array of `paramCount` pointers,
  corresponding to the arguments of `function`.

- `extraCount` is reserved for future use.

- `pExtras` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Kernel parameters of `function` are specified via `pParams`, very much
the same way as described in
[cuLaunchKernel](https://docs.nvidia.com/cuda/cuda-driver-api/group__CUDA__EXEC.html#group__CUDA__EXEC_1gb8f3dc3031b40da29d5f9a7139e52e15)

If `function` has N parameters, then `pParams` **must** be an array of N
pointers and `paramCount` **must** be set to N. Each of
`kernelParams`\[0\] through `kernelParams`\[N-1\] **must** point to a
region of memory from which the actual kernel parameter will be copied.
The number of kernel parameters and their offsets and sizes are not
specified here as that information is stored in the
[VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html) object.

The application-owned memory pointed to by `pParams` and
`kernelParams`\[0\] through `kernelParams`\[N-1\] are consumed
immediately, and **may** be altered or freed after
[vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCudaLaunchKernelNV.html) has returned.

Valid Usage

- <a href="#VUID-VkCudaLaunchInfoNV-gridDimX-09406"
  id="VUID-VkCudaLaunchInfoNV-gridDimX-09406"></a>
  VUID-VkCudaLaunchInfoNV-gridDimX-09406  
  `gridDimX` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\]

- <a href="#VUID-VkCudaLaunchInfoNV-gridDimY-09407"
  id="VUID-VkCudaLaunchInfoNV-gridDimY-09407"></a>
  VUID-VkCudaLaunchInfoNV-gridDimY-09407  
  `gridDimY` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\]

- <a href="#VUID-VkCudaLaunchInfoNV-gridDimZ-09408"
  id="VUID-VkCudaLaunchInfoNV-gridDimZ-09408"></a>
  VUID-VkCudaLaunchInfoNV-gridDimZ-09408  
  `gridDimZ` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\]

- <a href="#VUID-VkCudaLaunchInfoNV-paramCount-09409"
  id="VUID-VkCudaLaunchInfoNV-paramCount-09409"></a>
  VUID-VkCudaLaunchInfoNV-paramCount-09409  
  `paramCount` **must** be the total amount of parameters listed in the
  `pParams` table

- <a href="#VUID-VkCudaLaunchInfoNV-pParams-09410"
  id="VUID-VkCudaLaunchInfoNV-pParams-09410"></a>
  VUID-VkCudaLaunchInfoNV-pParams-09410  
  `pParams` **must** be a pointer to a table of `paramCount` parameters,
  corresponding to the arguments of `function`

- <a href="#VUID-VkCudaLaunchInfoNV-extraCount-09411"
  id="VUID-VkCudaLaunchInfoNV-extraCount-09411"></a>
  VUID-VkCudaLaunchInfoNV-extraCount-09411  
  `extraCount` must be 0

- <a href="#VUID-VkCudaLaunchInfoNV-pExtras-09412"
  id="VUID-VkCudaLaunchInfoNV-pExtras-09412"></a>
  VUID-VkCudaLaunchInfoNV-pExtras-09412  
  `pExtras` must be NULL

Valid Usage (Implicit)

- <a href="#VUID-VkCudaLaunchInfoNV-sType-sType"
  id="VUID-VkCudaLaunchInfoNV-sType-sType"></a>
  VUID-VkCudaLaunchInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CUDA_LAUNCH_INFO_NV`

- <a href="#VUID-VkCudaLaunchInfoNV-pNext-pNext"
  id="VUID-VkCudaLaunchInfoNV-pNext-pNext"></a>
  VUID-VkCudaLaunchInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCudaLaunchInfoNV-function-parameter"
  id="VUID-VkCudaLaunchInfoNV-function-parameter"></a>
  VUID-VkCudaLaunchInfoNV-function-parameter  
  `function` **must** be a valid
  [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaFunctionNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCudaLaunchKernelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCudaLaunchKernelNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCudaLaunchInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
